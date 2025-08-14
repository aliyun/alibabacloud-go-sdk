// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableUnderstandingResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetTableUnderstandingResultRequest
	GetId() *string
}

type GetTableUnderstandingResultRequest struct {
	// example:
	//
	// docmind-20220816-1e89d65c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTableUnderstandingResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableUnderstandingResultRequest) GoString() string {
	return s.String()
}

func (s *GetTableUnderstandingResultRequest) GetId() *string {
	return s.Id
}

func (s *GetTableUnderstandingResultRequest) SetId(v string) *GetTableUnderstandingResultRequest {
	s.Id = &v
	return s
}

func (s *GetTableUnderstandingResultRequest) Validate() error {
	return dara.Validate(s)
}
