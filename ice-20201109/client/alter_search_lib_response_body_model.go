// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterSearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AlterSearchLibResponseBody
	GetCode() *string
	SetRequestId(v string) *AlterSearchLibResponseBody
	GetRequestId() *string
	SetSearchLibName(v string) *AlterSearchLibResponseBody
	GetSearchLibName() *string
	SetSuccess(v string) *AlterSearchLibResponseBody
	GetSuccess() *string
}

type AlterSearchLibResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AlterSearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AlterSearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *AlterSearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *AlterSearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AlterSearchLibResponseBody) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *AlterSearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AlterSearchLibResponseBody) SetCode(v string) *AlterSearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *AlterSearchLibResponseBody) SetRequestId(v string) *AlterSearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *AlterSearchLibResponseBody) SetSearchLibName(v string) *AlterSearchLibResponseBody {
	s.SearchLibName = &v
	return s
}

func (s *AlterSearchLibResponseBody) SetSuccess(v string) *AlterSearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *AlterSearchLibResponseBody) Validate() error {
	return dara.Validate(s)
}
