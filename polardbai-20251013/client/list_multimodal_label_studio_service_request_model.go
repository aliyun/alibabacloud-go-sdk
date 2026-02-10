// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalLabelStudioServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListMultimodalLabelStudioServiceRequest
	GetDBClusterId() *string
}

type ListMultimodalLabelStudioServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s ListMultimodalLabelStudioServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalLabelStudioServiceRequest) GoString() string {
	return s.String()
}

func (s *ListMultimodalLabelStudioServiceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListMultimodalLabelStudioServiceRequest) SetDBClusterId(v string) *ListMultimodalLabelStudioServiceRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListMultimodalLabelStudioServiceRequest) Validate() error {
	return dara.Validate(s)
}
