// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportObjectSuffixResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *ListSupportObjectSuffixResponseBody
	GetData() []*string
	SetRequestId(v string) *ListSupportObjectSuffixResponseBody
	GetRequestId() *string
}

type ListSupportObjectSuffixResponseBody struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 2220FE66-76EF-5D9D-A0EE-3221CC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSupportObjectSuffixResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupportObjectSuffixResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupportObjectSuffixResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListSupportObjectSuffixResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupportObjectSuffixResponseBody) SetData(v []*string) *ListSupportObjectSuffixResponseBody {
	s.Data = v
	return s
}

func (s *ListSupportObjectSuffixResponseBody) SetRequestId(v string) *ListSupportObjectSuffixResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupportObjectSuffixResponseBody) Validate() error {
	return dara.Validate(s)
}
