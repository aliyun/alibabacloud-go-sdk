// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelKeyDeletionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *CancelKeyDeletionRequest
	GetKeyId() *string
}

type CancelKeyDeletionRequest struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s CancelKeyDeletionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelKeyDeletionRequest) GoString() string {
	return s.String()
}

func (s *CancelKeyDeletionRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *CancelKeyDeletionRequest) SetKeyId(v string) *CancelKeyDeletionRequest {
	s.KeyId = &v
	return s
}

func (s *CancelKeyDeletionRequest) Validate() error {
	return dara.Validate(s)
}
