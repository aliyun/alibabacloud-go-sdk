// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v string) *GetPublicKeyRequest
	GetDryRun() *string
	SetKeyId(v string) *GetPublicKeyRequest
	GetKeyId() *string
	SetKeyVersionId(v string) *GetPublicKeyRequest
	GetKeyVersionId() *string
}

type GetPublicKeyRequest struct {
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The globally unique ID of the CMK. You can also set this parameter to an alias that is bound to the CMK. For more information, see [Use aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s GetPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *GetPublicKeyRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *GetPublicKeyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GetPublicKeyRequest) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *GetPublicKeyRequest) SetDryRun(v string) *GetPublicKeyRequest {
	s.DryRun = &v
	return s
}

func (s *GetPublicKeyRequest) SetKeyId(v string) *GetPublicKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GetPublicKeyRequest) SetKeyVersionId(v string) *GetPublicKeyRequest {
	s.KeyVersionId = &v
	return s
}

func (s *GetPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
