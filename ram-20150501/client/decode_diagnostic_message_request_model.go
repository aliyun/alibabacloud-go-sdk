// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecodeDiagnosticMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncodedDiagnosticMessage(v string) *DecodeDiagnosticMessageRequest
	GetEncodedDiagnosticMessage() *string
}

type DecodeDiagnosticMessageRequest struct {
	// The encoded diagnostic information in the response that contains an access denied error. The error is caused by no RAM permissions.
	//
	// example:
	//
	// AQEAAAAAZBgxr0U1MjA1NTM1LUM4BBktMzE5RS1CODgxLUU1QTI0RDNFQTM1****
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
}

func (s DecodeDiagnosticMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s DecodeDiagnosticMessageRequest) GoString() string {
	return s.String()
}

func (s *DecodeDiagnosticMessageRequest) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DecodeDiagnosticMessageRequest) SetEncodedDiagnosticMessage(v string) *DecodeDiagnosticMessageRequest {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DecodeDiagnosticMessageRequest) Validate() error {
	return dara.Validate(s)
}
