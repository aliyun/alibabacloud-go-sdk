// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUdfArtifactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUdfArtifactName(v string) *GetUdfArtifactsRequest
	GetUdfArtifactName() *string
}

type GetUdfArtifactsRequest struct {
	// The name of the JAR or Python file that corresponds to the UDF.
	//
	// example:
	//
	// test-udf
	UdfArtifactName *string `json:"udfArtifactName,omitempty" xml:"udfArtifactName,omitempty"`
}

func (s GetUdfArtifactsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUdfArtifactsRequest) GoString() string {
	return s.String()
}

func (s *GetUdfArtifactsRequest) GetUdfArtifactName() *string {
	return s.UdfArtifactName
}

func (s *GetUdfArtifactsRequest) SetUdfArtifactName(v string) *GetUdfArtifactsRequest {
	s.UdfArtifactName = &v
	return s
}

func (s *GetUdfArtifactsRequest) Validate() error {
	return dara.Validate(s)
}
