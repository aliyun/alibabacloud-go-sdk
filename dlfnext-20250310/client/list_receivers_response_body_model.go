// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReceiversResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListReceiversResponseBody
	GetNextPageToken() *string
	SetReceivers(v []*Receiver) *ListReceiversResponseBody
	GetReceivers() []*Receiver
}

type ListReceiversResponseBody struct {
	// example:
	//
	// ""
	NextPageToken *string     `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Receivers     []*Receiver `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s ListReceiversResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListReceiversResponseBody) GoString() string {
	return s.String()
}

func (s *ListReceiversResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListReceiversResponseBody) GetReceivers() []*Receiver {
	return s.Receivers
}

func (s *ListReceiversResponseBody) SetNextPageToken(v string) *ListReceiversResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListReceiversResponseBody) SetReceivers(v []*Receiver) *ListReceiversResponseBody {
	s.Receivers = v
	return s
}

func (s *ListReceiversResponseBody) Validate() error {
	return dara.Validate(s)
}
