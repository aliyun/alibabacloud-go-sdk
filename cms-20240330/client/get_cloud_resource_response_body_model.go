// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetCloudResourceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetCloudResourceResponseBody
	GetRequestId() *string
}

type GetCloudResourceResponseBody struct {
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCloudResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudResourceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCloudResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudResourceResponseBody) SetRegionId(v string) *GetCloudResourceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetCloudResourceResponseBody) SetRequestId(v string) *GetCloudResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
