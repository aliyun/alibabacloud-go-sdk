// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvidedSharesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListProvidedSharesResponseBody
	GetNextPageToken() *string
	SetShares(v []*Share) *ListProvidedSharesResponseBody
	GetShares() []*Share
}

type ListProvidedSharesResponseBody struct {
	// example:
	//
	// ""
	NextPageToken *string  `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Shares        []*Share `json:"shares,omitempty" xml:"shares,omitempty" type:"Repeated"`
}

func (s ListProvidedSharesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProvidedSharesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProvidedSharesResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListProvidedSharesResponseBody) GetShares() []*Share {
	return s.Shares
}

func (s *ListProvidedSharesResponseBody) SetNextPageToken(v string) *ListProvidedSharesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListProvidedSharesResponseBody) SetShares(v []*Share) *ListProvidedSharesResponseBody {
	s.Shares = v
	return s
}

func (s *ListProvidedSharesResponseBody) Validate() error {
	return dara.Validate(s)
}
