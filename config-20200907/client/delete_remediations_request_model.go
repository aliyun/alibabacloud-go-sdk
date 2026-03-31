// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRemediationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemediationIds(v string) *DeleteRemediationsRequest
	GetRemediationIds() *string
}

type DeleteRemediationsRequest struct {
	// The ID of the remediation template. Separate multiple remediation template IDs with commas (,).
	//
	// For more information about how to obtain the ID of a remediation template, see [ListRemediations](https://help.aliyun.com/document_detail/270772.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-909ba2d4716700eb****
	RemediationIds *string `json:"RemediationIds,omitempty" xml:"RemediationIds,omitempty"`
}

func (s DeleteRemediationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRemediationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteRemediationsRequest) GetRemediationIds() *string {
	return s.RemediationIds
}

func (s *DeleteRemediationsRequest) SetRemediationIds(v string) *DeleteRemediationsRequest {
	s.RemediationIds = &v
	return s
}

func (s *DeleteRemediationsRequest) Validate() error {
	return dara.Validate(s)
}
