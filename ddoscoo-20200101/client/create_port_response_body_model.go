// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePortResponseBody
	GetRequestId() *string
}

type CreatePortResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4787A9A6-8230-4B4A-8211-AFBF7C416B4D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePortResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePortResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePortResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePortResponseBody) SetRequestId(v string) *CreatePortResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePortResponseBody) Validate() error {
	return dara.Validate(s)
}
