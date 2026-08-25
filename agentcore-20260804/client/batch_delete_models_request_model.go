// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *BatchDeleteModelsRequestBody) *BatchDeleteModelsRequest
	GetBody() *BatchDeleteModelsRequestBody
	SetClientToken(v string) *BatchDeleteModelsRequest
	GetClientToken() *string
}

type BatchDeleteModelsRequest struct {
	Body *BatchDeleteModelsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s BatchDeleteModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteModelsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelsRequest) GetBody() *BatchDeleteModelsRequestBody {
	return s.Body
}

func (s *BatchDeleteModelsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *BatchDeleteModelsRequest) SetBody(v *BatchDeleteModelsRequestBody) *BatchDeleteModelsRequest {
	s.Body = v
	return s
}

func (s *BatchDeleteModelsRequest) SetClientToken(v string) *BatchDeleteModelsRequest {
	s.ClientToken = &v
	return s
}

func (s *BatchDeleteModelsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchDeleteModelsRequestBody struct {
	// This parameter is required.
	ModelIds []*string `json:"modelIds,omitempty" xml:"modelIds,omitempty" type:"Repeated"`
}

func (s BatchDeleteModelsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteModelsRequestBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelsRequestBody) GetModelIds() []*string {
	return s.ModelIds
}

func (s *BatchDeleteModelsRequestBody) SetModelIds(v []*string) *BatchDeleteModelsRequestBody {
	s.ModelIds = v
	return s
}

func (s *BatchDeleteModelsRequestBody) Validate() error {
	return dara.Validate(s)
}
