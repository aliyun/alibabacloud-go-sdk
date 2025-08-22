// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDcdnProjectExistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *CheckDcdnProjectExistResponseBodyContent) *CheckDcdnProjectExistResponseBody
	GetContent() *CheckDcdnProjectExistResponseBodyContent
	SetRequestId(v string) *CheckDcdnProjectExistResponseBody
	GetRequestId() *string
}

type CheckDcdnProjectExistResponseBody struct {
	// The returned results.
	Content *CheckDcdnProjectExistResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// b021e538-9dde-46ed-a1f2-9469da8f3e77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDcdnProjectExistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDcdnProjectExistResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDcdnProjectExistResponseBody) GetContent() *CheckDcdnProjectExistResponseBodyContent {
	return s.Content
}

func (s *CheckDcdnProjectExistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDcdnProjectExistResponseBody) SetContent(v *CheckDcdnProjectExistResponseBodyContent) *CheckDcdnProjectExistResponseBody {
	s.Content = v
	return s
}

func (s *CheckDcdnProjectExistResponseBody) SetRequestId(v string) *CheckDcdnProjectExistResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDcdnProjectExistResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckDcdnProjectExistResponseBodyContent struct {
	// Indicates whether the real-time log delivery project exists. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	Exist *string `json:"Exist,omitempty" xml:"Exist,omitempty"`
}

func (s CheckDcdnProjectExistResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CheckDcdnProjectExistResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CheckDcdnProjectExistResponseBodyContent) GetExist() *string {
	return s.Exist
}

func (s *CheckDcdnProjectExistResponseBodyContent) SetExist(v string) *CheckDcdnProjectExistResponseBodyContent {
	s.Exist = &v
	return s
}

func (s *CheckDcdnProjectExistResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
