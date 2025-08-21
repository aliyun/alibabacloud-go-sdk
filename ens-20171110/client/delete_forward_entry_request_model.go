// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardEntryId(v string) *DeleteForwardEntryRequest
	GetForwardEntryId() *string
}

type DeleteForwardEntryRequest struct {
	// The ID of the DNAT entry that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// fwd-5tfk8qgepr9ijjkqxt8do****
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
}

func (s DeleteForwardEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteForwardEntryRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DeleteForwardEntryRequest) SetForwardEntryId(v string) *DeleteForwardEntryRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *DeleteForwardEntryRequest) Validate() error {
	return dara.Validate(s)
}
