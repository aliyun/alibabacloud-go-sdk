// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceVersionResponseBody
	GetRequestId() *string
}

type UpdateServiceVersionResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Service [foo] in region [cn-shanghai] is starting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceVersionResponseBody) SetMessage(v string) *UpdateServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceVersionResponseBody) SetRequestId(v string) *UpdateServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
