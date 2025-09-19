// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTrailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceTrailResponseBody
	GetRequestId() *string
	SetServiceTrail(v *GetServiceTrailResponseBodyServiceTrail) *GetServiceTrailResponseBody
	GetServiceTrail() *GetServiceTrailResponseBodyServiceTrail
}

type GetServiceTrailResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C699E4E4-F2F4-58FC-A949-457FFE59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the service trail.
	ServiceTrail *GetServiceTrailResponseBodyServiceTrail `json:"ServiceTrail,omitempty" xml:"ServiceTrail,omitempty" type:"Struct"`
}

func (s GetServiceTrailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTrailResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTrailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceTrailResponseBody) GetServiceTrail() *GetServiceTrailResponseBodyServiceTrail {
	return s.ServiceTrail
}

func (s *GetServiceTrailResponseBody) SetRequestId(v string) *GetServiceTrailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceTrailResponseBody) SetServiceTrail(v *GetServiceTrailResponseBodyServiceTrail) *GetServiceTrailResponseBody {
	s.ServiceTrail = v
	return s
}

func (s *GetServiceTrailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceTrailResponseBodyServiceTrail struct {
	// The status of the service trail. Valid values:
	//
	// 	- **on:**
	//
	// 	- **off:**
	//
	// example:
	//
	// on
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The timestamp generated when the service trail was created. Unit: milliseconds.
	//
	// example:
	//
	// 1687250241000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp generated when the service trail was last updated. Unit: milliseconds.
	//
	// example:
	//
	// 1687250241000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetServiceTrailResponseBodyServiceTrail) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTrailResponseBodyServiceTrail) GoString() string {
	return s.String()
}

func (s *GetServiceTrailResponseBodyServiceTrail) GetConfig() *string {
	return s.Config
}

func (s *GetServiceTrailResponseBodyServiceTrail) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetServiceTrailResponseBodyServiceTrail) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetServiceTrailResponseBodyServiceTrail) SetConfig(v string) *GetServiceTrailResponseBodyServiceTrail {
	s.Config = &v
	return s
}

func (s *GetServiceTrailResponseBodyServiceTrail) SetCreateTime(v int64) *GetServiceTrailResponseBodyServiceTrail {
	s.CreateTime = &v
	return s
}

func (s *GetServiceTrailResponseBodyServiceTrail) SetUpdateTime(v int64) *GetServiceTrailResponseBodyServiceTrail {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceTrailResponseBodyServiceTrail) Validate() error {
	return dara.Validate(s)
}
