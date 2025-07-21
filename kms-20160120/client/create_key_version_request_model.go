// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *CreateKeyVersionRequest
	GetKeyId() *string
}

type CreateKeyVersionRequest struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  You can also set the value to an alias that is bound to the CMK. For more information, see [Overview of aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0b30658a-ed1a-4922-b8f7-a673ca9c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s CreateKeyVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *CreateKeyVersionRequest) SetKeyId(v string) *CreateKeyVersionRequest {
	s.KeyId = &v
	return s
}

func (s *CreateKeyVersionRequest) Validate() error {
	return dara.Validate(s)
}
