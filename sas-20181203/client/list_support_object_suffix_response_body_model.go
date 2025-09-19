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
	// The returned data.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 571B2642-BF51-5BDD-906B-D2340DB9****
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
