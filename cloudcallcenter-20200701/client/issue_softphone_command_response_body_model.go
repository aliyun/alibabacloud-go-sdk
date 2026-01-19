// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIssueSoftphoneCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IssueSoftphoneCommandResponseBody
	GetCode() *string
	SetData(v string) *IssueSoftphoneCommandResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *IssueSoftphoneCommandResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *IssueSoftphoneCommandResponseBody
	GetMessage() *string
	SetRequestId(v string) *IssueSoftphoneCommandResponseBody
	GetRequestId() *string
}

type IssueSoftphoneCommandResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IssueSoftphoneCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IssueSoftphoneCommandResponseBody) GoString() string {
	return s.String()
}

func (s *IssueSoftphoneCommandResponseBody) GetCode() *string {
	return s.Code
}

func (s *IssueSoftphoneCommandResponseBody) GetData() *string {
	return s.Data
}

func (s *IssueSoftphoneCommandResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *IssueSoftphoneCommandResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IssueSoftphoneCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IssueSoftphoneCommandResponseBody) SetCode(v string) *IssueSoftphoneCommandResponseBody {
	s.Code = &v
	return s
}

func (s *IssueSoftphoneCommandResponseBody) SetData(v string) *IssueSoftphoneCommandResponseBody {
	s.Data = &v
	return s
}

func (s *IssueSoftphoneCommandResponseBody) SetHttpStatusCode(v int32) *IssueSoftphoneCommandResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *IssueSoftphoneCommandResponseBody) SetMessage(v string) *IssueSoftphoneCommandResponseBody {
	s.Message = &v
	return s
}

func (s *IssueSoftphoneCommandResponseBody) SetRequestId(v string) *IssueSoftphoneCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *IssueSoftphoneCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
