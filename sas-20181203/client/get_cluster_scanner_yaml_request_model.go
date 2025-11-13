// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterScannerYamlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetClusterScannerYamlRequest
	GetClusterId() *string
}

type GetClusterScannerYamlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd49575861a3044d49c954e4b3911****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetClusterScannerYamlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterScannerYamlRequest) GoString() string {
	return s.String()
}

func (s *GetClusterScannerYamlRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterScannerYamlRequest) SetClusterId(v string) *GetClusterScannerYamlRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterScannerYamlRequest) Validate() error {
	return dara.Validate(s)
}
