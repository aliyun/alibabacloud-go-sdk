// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIDCImportCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateIDCImportCommandResponseBody
	GetCode() *string
	SetData(v string) *CreateIDCImportCommandResponseBody
	GetData() *string
	SetMessage(v string) *CreateIDCImportCommandResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIDCImportCommandResponseBody
	GetRequestId() *string
}

type CreateIDCImportCommandResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The generated import command.
	//
	// example:
	//
	// wget -q -O /tmp/install.sh http://edas-hz.oss-cn-hangzhou-internal.aliyuncs.com/install.sh && sh /tmp/install.sh -idcToken xxxx-xxxxx-xxxxx-xxxxxxx -edasId xxxxxxxxxxxxxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// b197-40ab-9155-7ca7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIDCImportCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIDCImportCommandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIDCImportCommandResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateIDCImportCommandResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateIDCImportCommandResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIDCImportCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIDCImportCommandResponseBody) SetCode(v string) *CreateIDCImportCommandResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIDCImportCommandResponseBody) SetData(v string) *CreateIDCImportCommandResponseBody {
	s.Data = &v
	return s
}

func (s *CreateIDCImportCommandResponseBody) SetMessage(v string) *CreateIDCImportCommandResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIDCImportCommandResponseBody) SetRequestId(v string) *CreateIDCImportCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIDCImportCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
