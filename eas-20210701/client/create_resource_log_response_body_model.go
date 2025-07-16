// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateResourceLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateResourceLogResponseBody
	GetRequestId() *string
}

type CreateResourceLogResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Resource [eas-r-asdasdasd] is deleted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateResourceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceLogResponseBody) SetMessage(v string) *CreateResourceLogResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResourceLogResponseBody) SetRequestId(v string) *CreateResourceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceLogResponseBody) Validate() error {
	return dara.Validate(s)
}
