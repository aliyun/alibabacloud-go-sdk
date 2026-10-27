// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateMemoryStoreRequestBody) *UpdateMemoryStoreRequest
	GetBody() *UpdateMemoryStoreRequestBody
	SetClientToken(v string) *UpdateMemoryStoreRequest
	GetClientToken() *string
}

type UpdateMemoryStoreRequest struct {
	// The update MemoryStore request.
	//
	// This parameter is required.
	Body *UpdateMemoryStoreRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client request token.
	//
	// example:
	//
	// memory-store-update-20260909-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateMemoryStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryStoreRequest) GetBody() *UpdateMemoryStoreRequestBody {
	return s.Body
}

func (s *UpdateMemoryStoreRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateMemoryStoreRequest) SetBody(v *UpdateMemoryStoreRequestBody) *UpdateMemoryStoreRequest {
	s.Body = v
	return s
}

func (s *UpdateMemoryStoreRequest) SetClientToken(v string) *UpdateMemoryStoreRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateMemoryStoreRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryStoreRequestBody struct {
	// The MemoryStore description, up to 1,024 characters.
	//
	// example:
	//
	// Memory store used to save long-term customer preferences
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The custom memory extraction instructions, up to 4,096 characters.
	//
	// example:
	//
	// Extract only long-term preferences explicitly expressed by the user
	ExtractInstructions *string `json:"extractInstructions,omitempty" xml:"extractInstructions,omitempty"`
}

func (s UpdateMemoryStoreRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryStoreRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryStoreRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateMemoryStoreRequestBody) GetExtractInstructions() *string {
	return s.ExtractInstructions
}

func (s *UpdateMemoryStoreRequestBody) SetDescription(v string) *UpdateMemoryStoreRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateMemoryStoreRequestBody) SetExtractInstructions(v string) *UpdateMemoryStoreRequestBody {
	s.ExtractInstructions = &v
	return s
}

func (s *UpdateMemoryStoreRequestBody) Validate() error {
	return dara.Validate(s)
}
