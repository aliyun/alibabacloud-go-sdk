// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateMergedTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Schema) *GenerateMergedTableRequest
	GetBody() *Schema
	SetSpec(v string) *GenerateMergedTableRequest
	GetSpec() *string
}

type GenerateMergedTableRequest struct {
	// The request body parameters.
	Body *Schema `json:"body,omitempty" xml:"body,omitempty"`
	// The specifications of the OpenSearch instance. This parameter is used to check whether the OpenSearch instance meets the special limits on an exclusive instance.
	//
	// Default value: opensearch.share.common.
	//
	// For more information, see the description of the spec field in the Quota topic.
	//
	// example:
	//
	// "opensearch.share.common"
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s GenerateMergedTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateMergedTableRequest) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableRequest) GetBody() *Schema {
	return s.Body
}

func (s *GenerateMergedTableRequest) GetSpec() *string {
	return s.Spec
}

func (s *GenerateMergedTableRequest) SetBody(v *Schema) *GenerateMergedTableRequest {
	s.Body = v
	return s
}

func (s *GenerateMergedTableRequest) SetSpec(v string) *GenerateMergedTableRequest {
	s.Spec = &v
	return s
}

func (s *GenerateMergedTableRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
