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
	// The reserved query request identifier. The current console backend does not read this parameter and does not provide idempotency guarantees for deduplication or recovery of the original processing result based on this identifier.
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
