// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRolloverDataStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RolloverDataStreamRequest
	GetClientToken() *string
}

type RolloverDataStreamRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RolloverDataStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s RolloverDataStreamRequest) GoString() string {
	return s.String()
}

func (s *RolloverDataStreamRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RolloverDataStreamRequest) SetClientToken(v string) *RolloverDataStreamRequest {
	s.ClientToken = &v
	return s
}

func (s *RolloverDataStreamRequest) Validate() error {
	return dara.Validate(s)
}
