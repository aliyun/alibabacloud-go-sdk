// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDevelopmentModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *GetDevelopmentModeResponseBody
	GetEnable() *string
	SetRequestId(v string) *GetDevelopmentModeResponseBody
	GetRequestId() *string
}

type GetDevelopmentModeResponseBody struct {
	Enable    *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDevelopmentModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDevelopmentModeResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevelopmentModeResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *GetDevelopmentModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDevelopmentModeResponseBody) SetEnable(v string) *GetDevelopmentModeResponseBody {
	s.Enable = &v
	return s
}

func (s *GetDevelopmentModeResponseBody) SetRequestId(v string) *GetDevelopmentModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevelopmentModeResponseBody) Validate() error {
	return dara.Validate(s)
}
