// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachWhitelistTemplateToInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DetachWhitelistTemplateToInstanceResponseBody
	GetCode() *string
	SetData(v *DetachWhitelistTemplateToInstanceResponseBodyData) *DetachWhitelistTemplateToInstanceResponseBody
	GetData() *DetachWhitelistTemplateToInstanceResponseBodyData
	SetHttpStatusCode(v int32) *DetachWhitelistTemplateToInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DetachWhitelistTemplateToInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetachWhitelistTemplateToInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetachWhitelistTemplateToInstanceResponseBody
	GetSuccess() *bool
}

type DetachWhitelistTemplateToInstanceResponseBody struct {
	Code           *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DetachWhitelistTemplateToInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachWhitelistTemplateToInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachWhitelistTemplateToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) GetData() *DetachWhitelistTemplateToInstanceResponseBodyData {
	return s.Data
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) SetCode(v string) *DetachWhitelistTemplateToInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) SetData(v *DetachWhitelistTemplateToInstanceResponseBodyData) *DetachWhitelistTemplateToInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) SetHttpStatusCode(v int32) *DetachWhitelistTemplateToInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) SetMessage(v string) *DetachWhitelistTemplateToInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) SetRequestId(v string) *DetachWhitelistTemplateToInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) SetSuccess(v bool) *DetachWhitelistTemplateToInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachWhitelistTemplateToInstanceResponseBodyData struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DetachWhitelistTemplateToInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetachWhitelistTemplateToInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyData) SetStatus(v string) *DetachWhitelistTemplateToInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
