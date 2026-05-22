// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSeoBypassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *GetSeoBypassResponseBody
	GetEnable() *string
	SetRequestId(v string) *GetSeoBypassResponseBody
	GetRequestId() *string
}

type GetSeoBypassResponseBody struct {
	// The status of the feature. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSeoBypassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSeoBypassResponseBody) GoString() string {
	return s.String()
}

func (s *GetSeoBypassResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *GetSeoBypassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSeoBypassResponseBody) SetEnable(v string) *GetSeoBypassResponseBody {
	s.Enable = &v
	return s
}

func (s *GetSeoBypassResponseBody) SetRequestId(v string) *GetSeoBypassResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSeoBypassResponseBody) Validate() error {
	return dara.Validate(s)
}
