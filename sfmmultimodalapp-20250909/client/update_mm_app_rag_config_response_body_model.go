// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRagConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMmAppRagConfigResponseBody
	GetCode() *string
	SetData(v *UpdateMmAppRagConfigResponseBodyData) *UpdateMmAppRagConfigResponseBody
	GetData() *UpdateMmAppRagConfigResponseBodyData
	SetHttpStatusCode(v int32) *UpdateMmAppRagConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMmAppRagConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMmAppRagConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMmAppRagConfigResponseBody
	GetSuccess() *bool
}

type UpdateMmAppRagConfigResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateMmAppRagConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppRagConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMmAppRagConfigResponseBody) GetData() *UpdateMmAppRagConfigResponseBodyData {
	return s.Data
}

func (s *UpdateMmAppRagConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMmAppRagConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMmAppRagConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmAppRagConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppRagConfigResponseBody) SetCode(v string) *UpdateMmAppRagConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMmAppRagConfigResponseBody) SetData(v *UpdateMmAppRagConfigResponseBodyData) *UpdateMmAppRagConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMmAppRagConfigResponseBody) SetHttpStatusCode(v int32) *UpdateMmAppRagConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMmAppRagConfigResponseBody) SetMessage(v string) *UpdateMmAppRagConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMmAppRagConfigResponseBody) SetRequestId(v string) *UpdateMmAppRagConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmAppRagConfigResponseBody) SetSuccess(v bool) *UpdateMmAppRagConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMmAppRagConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMmAppRagConfigResponseBodyData struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMmAppRagConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagConfigResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMmAppRagConfigResponseBodyData) SetSuccess(v bool) *UpdateMmAppRagConfigResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateMmAppRagConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
