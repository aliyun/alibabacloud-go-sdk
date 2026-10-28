// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOncallContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactIds(v []*string) *ListOncallContactsResponseBody
	GetContactIds() []*string
	SetRequestId(v string) *ListOncallContactsResponseBody
	GetRequestId() *string
}

type ListOncallContactsResponseBody struct {
	// The IDs of the contacts.
	ContactIds []*string `json:"contactIds,omitempty" xml:"contactIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListOncallContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOncallContactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOncallContactsResponseBody) GetContactIds() []*string {
	return s.ContactIds
}

func (s *ListOncallContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOncallContactsResponseBody) SetContactIds(v []*string) *ListOncallContactsResponseBody {
	s.ContactIds = v
	return s
}

func (s *ListOncallContactsResponseBody) SetRequestId(v string) *ListOncallContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOncallContactsResponseBody) Validate() error {
	return dara.Validate(s)
}
