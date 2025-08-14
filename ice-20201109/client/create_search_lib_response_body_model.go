// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSearchLibResponseBody
	GetCode() *string
	SetRequestId(v string) *CreateSearchLibResponseBody
	GetRequestId() *string
	SetSearchLibName(v string) *CreateSearchLibResponseBody
	GetSearchLibName() *string
	SetSuccess(v string) *CreateSearchLibResponseBody
	GetSuccess() *string
}

type CreateSearchLibResponseBody struct {
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
	// The name of the search library.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSearchLibResponseBody) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *CreateSearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateSearchLibResponseBody) SetCode(v string) *CreateSearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSearchLibResponseBody) SetRequestId(v string) *CreateSearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSearchLibResponseBody) SetSearchLibName(v string) *CreateSearchLibResponseBody {
	s.SearchLibName = &v
	return s
}

func (s *CreateSearchLibResponseBody) SetSuccess(v string) *CreateSearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSearchLibResponseBody) Validate() error {
	return dara.Validate(s)
}
