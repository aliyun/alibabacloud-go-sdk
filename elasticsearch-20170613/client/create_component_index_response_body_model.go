// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateComponentIndexResponseBody
	GetRequestId() *string
	SetResult(v bool) *CreateComponentIndexResponseBody
	GetResult() *bool
}

type CreateComponentIndexResponseBody struct {
	// example:
	//
	// C20022BA-5382-4339-89FB-30AF48A05431
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateComponentIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComponentIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComponentIndexResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CreateComponentIndexResponseBody) SetRequestId(v string) *CreateComponentIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComponentIndexResponseBody) SetResult(v bool) *CreateComponentIndexResponseBody {
	s.Result = &v
	return s
}

func (s *CreateComponentIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
