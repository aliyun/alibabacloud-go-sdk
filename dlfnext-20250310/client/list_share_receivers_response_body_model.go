// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShareReceiversResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListShareReceiversResponseBody
	GetNextPageToken() *string
	SetReceivers(v []*Receiver) *ListShareReceiversResponseBody
	GetReceivers() []*Receiver
}

type ListShareReceiversResponseBody struct {
	// example:
	//
	// “”
	NextPageToken *string     `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Receivers     []*Receiver `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
}

func (s ListShareReceiversResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListShareReceiversResponseBody) GoString() string {
	return s.String()
}

func (s *ListShareReceiversResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListShareReceiversResponseBody) GetReceivers() []*Receiver {
	return s.Receivers
}

func (s *ListShareReceiversResponseBody) SetNextPageToken(v string) *ListShareReceiversResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListShareReceiversResponseBody) SetReceivers(v []*Receiver) *ListShareReceiversResponseBody {
	s.Receivers = v
	return s
}

func (s *ListShareReceiversResponseBody) Validate() error {
	if s.Receivers != nil {
		for _, item := range s.Receivers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
