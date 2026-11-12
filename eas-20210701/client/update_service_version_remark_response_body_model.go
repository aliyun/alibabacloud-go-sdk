// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceVersionRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceVersionRemarkResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceVersionRemarkResponseBody
	GetRequestId() *string
}

type UpdateServiceVersionRemarkResponseBody struct {
	// The message returned in the response.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceVersionRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceVersionRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionRemarkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceVersionRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceVersionRemarkResponseBody) SetMessage(v string) *UpdateServiceVersionRemarkResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceVersionRemarkResponseBody) SetRequestId(v string) *UpdateServiceVersionRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceVersionRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
