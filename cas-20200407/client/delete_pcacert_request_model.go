// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePCACertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *DeletePCACertRequest
	GetIdentifier() *string
}

type DeletePCACertRequest struct {
	// The unique identifier of the certificate. You can call the [ListCert](https://help.aliyun.com/document_detail/452331.html) operation to query the unique identifiers of certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccaf0c629c2be1e2ab
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s DeletePCACertRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePCACertRequest) GoString() string {
	return s.String()
}

func (s *DeletePCACertRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DeletePCACertRequest) SetIdentifier(v string) *DeletePCACertRequest {
	s.Identifier = &v
	return s
}

func (s *DeletePCACertRequest) Validate() error {
	return dara.Validate(s)
}
