// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceMigrationAvailableNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplaceMigrationAvailableNumbersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ReplaceMigrationAvailableNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ReplaceMigrationAvailableNumbersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReplaceMigrationAvailableNumbersResponseBody
	GetRequestId() *string
}

type ReplaceMigrationAvailableNumbersResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceMigrationAvailableNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceMigrationAvailableNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) SetCode(v string) *ReplaceMigrationAvailableNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) SetHttpStatusCode(v int32) *ReplaceMigrationAvailableNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) SetMessage(v string) *ReplaceMigrationAvailableNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) SetRequestId(v string) *ReplaceMigrationAvailableNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}
