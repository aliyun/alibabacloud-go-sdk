// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModuleId(v string) *CreateModuleResponseBody
	GetModuleId() *string
	SetRequestId(v string) *CreateModuleResponseBody
	GetRequestId() *string
}

type CreateModuleResponseBody struct {
	// example:
	//
	// mod-518855d9a058cfffcc446d8fe3c99
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// 0D797DC3-FF04-5C21-81EB-92C7799512E3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModuleResponseBody) GetModuleId() *string {
	return s.ModuleId
}

func (s *CreateModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModuleResponseBody) SetModuleId(v string) *CreateModuleResponseBody {
	s.ModuleId = &v
	return s
}

func (s *CreateModuleResponseBody) SetRequestId(v string) *CreateModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModuleResponseBody) Validate() error {
	return dara.Validate(s)
}
