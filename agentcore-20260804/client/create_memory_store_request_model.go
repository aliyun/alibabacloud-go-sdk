// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateMemoryStoreRequestBody) *CreateMemoryStoreRequest
	GetBody() *CreateMemoryStoreRequestBody
	SetClientToken(v string) *CreateMemoryStoreRequest
	GetClientToken() *string
}

type CreateMemoryStoreRequest struct {
	// The MemoryStore creation request.
	//
	// This parameter is required.
	Body *CreateMemoryStoreRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token.
	//
	// example:
	//
	// memory-store-create-20260909-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateMemoryStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryStoreRequest) GetBody() *CreateMemoryStoreRequestBody {
	return s.Body
}

func (s *CreateMemoryStoreRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateMemoryStoreRequest) SetBody(v *CreateMemoryStoreRequestBody) *CreateMemoryStoreRequest {
	s.Body = v
	return s
}

func (s *CreateMemoryStoreRequest) SetClientToken(v string) *CreateMemoryStoreRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMemoryStoreRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryStoreRequestBody struct {
	// The description of the MemoryStore. Maximum length: 1,024 characters.
	//
	// example:
	//
	// Used to store long-term user preferences
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The custom extraction instructions. Maximum length: 4,096 characters.
	//
	// example:
	//
	// Extract only explicitly expressed long-term user preferences
	ExtractInstructions *string `json:"extractInstructions,omitempty" xml:"extractInstructions,omitempty"`
	// The name of the MemoryStore. The name must be 1 to 64 characters in length and can contain only lowercase letters, digits, and hyphens (-). The name cannot be modified after creation.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer-memory
	MemoryStoreName *string `json:"memoryStoreName,omitempty" xml:"memoryStoreName,omitempty"`
}

func (s CreateMemoryStoreRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryStoreRequestBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryStoreRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateMemoryStoreRequestBody) GetExtractInstructions() *string {
	return s.ExtractInstructions
}

func (s *CreateMemoryStoreRequestBody) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *CreateMemoryStoreRequestBody) SetDescription(v string) *CreateMemoryStoreRequestBody {
	s.Description = &v
	return s
}

func (s *CreateMemoryStoreRequestBody) SetExtractInstructions(v string) *CreateMemoryStoreRequestBody {
	s.ExtractInstructions = &v
	return s
}

func (s *CreateMemoryStoreRequestBody) SetMemoryStoreName(v string) *CreateMemoryStoreRequestBody {
	s.MemoryStoreName = &v
	return s
}

func (s *CreateMemoryStoreRequestBody) Validate() error {
	return dara.Validate(s)
}
