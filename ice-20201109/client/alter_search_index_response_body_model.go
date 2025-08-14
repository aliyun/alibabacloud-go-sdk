// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterSearchIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AlterSearchIndexResponseBody
	GetCode() *string
	SetRequestId(v string) *AlterSearchIndexResponseBody
	GetRequestId() *string
	SetSuccess(v string) *AlterSearchIndexResponseBody
	GetSuccess() *string
}

type AlterSearchIndexResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AlterSearchIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AlterSearchIndexResponseBody) GoString() string {
	return s.String()
}

func (s *AlterSearchIndexResponseBody) GetCode() *string {
	return s.Code
}

func (s *AlterSearchIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AlterSearchIndexResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AlterSearchIndexResponseBody) SetCode(v string) *AlterSearchIndexResponseBody {
	s.Code = &v
	return s
}

func (s *AlterSearchIndexResponseBody) SetRequestId(v string) *AlterSearchIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *AlterSearchIndexResponseBody) SetSuccess(v string) *AlterSearchIndexResponseBody {
	s.Success = &v
	return s
}

func (s *AlterSearchIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
