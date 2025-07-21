// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKmsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKmsInstanceId(v string) *GetKmsInstanceRequest
	GetKmsInstanceId() *string
}

type GetKmsInstanceRequest struct {
	// The ID of the KMS instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// kst-bjj62f5ba3dnpb6v8****
	KmsInstanceId *string `json:"KmsInstanceId,omitempty" xml:"KmsInstanceId,omitempty"`
}

func (s GetKmsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKmsInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceRequest) GetKmsInstanceId() *string {
	return s.KmsInstanceId
}

func (s *GetKmsInstanceRequest) SetKmsInstanceId(v string) *GetKmsInstanceRequest {
	s.KmsInstanceId = &v
	return s
}

func (s *GetKmsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
