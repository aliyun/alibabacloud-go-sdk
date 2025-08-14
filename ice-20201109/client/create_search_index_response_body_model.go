// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSearchIndexResponseBody
	GetCode() *string
	SetRequestId(v string) *CreateSearchIndexResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateSearchIndexResponseBody
	GetSuccess() *string
}

type CreateSearchIndexResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// *****ACB-44F2-5F2D-88D7-1283E70*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSearchIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSearchIndexResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSearchIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSearchIndexResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateSearchIndexResponseBody) SetCode(v string) *CreateSearchIndexResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSearchIndexResponseBody) SetRequestId(v string) *CreateSearchIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSearchIndexResponseBody) SetSuccess(v string) *CreateSearchIndexResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSearchIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
