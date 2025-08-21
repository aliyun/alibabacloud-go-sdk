// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDataStreamRequest
	GetClientToken() *string
}

type DeleteDataStreamRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteDataStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataStreamRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataStreamRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDataStreamRequest) SetClientToken(v string) *DeleteDataStreamRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDataStreamRequest) Validate() error {
	return dara.Validate(s)
}
