// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertMcpToFreeEditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ConvertMcpToFreeEditRequest
	GetClientToken() *string
}

type ConvertMcpToFreeEditRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ConvertMcpToFreeEditRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertMcpToFreeEditRequest) GoString() string {
	return s.String()
}

func (s *ConvertMcpToFreeEditRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ConvertMcpToFreeEditRequest) SetClientToken(v string) *ConvertMcpToFreeEditRequest {
	s.ClientToken = &v
	return s
}

func (s *ConvertMcpToFreeEditRequest) Validate() error {
	return dara.Validate(s)
}
