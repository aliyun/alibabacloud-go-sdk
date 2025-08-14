// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDocParserStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *QueryDocParserStatusRequest
	GetId() *string
}

type QueryDocParserStatusRequest struct {
	// example:
	//
	// docmind-20220816-1e89d65c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryDocParserStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDocParserStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDocParserStatusRequest) GetId() *string {
	return s.Id
}

func (s *QueryDocParserStatusRequest) SetId(v string) *QueryDocParserStatusRequest {
	s.Id = &v
	return s
}

func (s *QueryDocParserStatusRequest) Validate() error {
	return dara.Validate(s)
}
