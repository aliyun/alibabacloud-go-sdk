// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareImageFacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *CompareImageFacesRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *CompareImageFacesRequest
	GetProjectName() *string
	SetSource(v *CompareImageFacesRequestSource) *CompareImageFacesRequest
	GetSource() *CompareImageFacesRequestSource
}

type CompareImageFacesRequest struct {
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The URLs of the two images for compression.
	Source *CompareImageFacesRequestSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
}

func (s CompareImageFacesRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareImageFacesRequest) GoString() string {
	return s.String()
}

func (s *CompareImageFacesRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CompareImageFacesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CompareImageFacesRequest) GetSource() *CompareImageFacesRequestSource {
	return s.Source
}

func (s *CompareImageFacesRequest) SetCredentialConfig(v *CredentialConfig) *CompareImageFacesRequest {
	s.CredentialConfig = v
	return s
}

func (s *CompareImageFacesRequest) SetProjectName(v string) *CompareImageFacesRequest {
	s.ProjectName = &v
	return s
}

func (s *CompareImageFacesRequest) SetSource(v *CompareImageFacesRequestSource) *CompareImageFacesRequest {
	s.Source = v
	return s
}

func (s *CompareImageFacesRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompareImageFacesRequestSource struct {
	// The OSS URL of the image file.
	//
	// Specify the URL in the `oss://<bucket>/<object>` format. `<bucket>` specifies the name of the OSS bucket that is in the same region as the current project. `<object>` specifies path of the object with the extension included.
	//
	// example:
	//
	// oss://test-bucket/test-object1
	URI1 *string `json:"URI1,omitempty" xml:"URI1,omitempty"`
	// The OSS URL of the image file.
	//
	// Specify the URL in the `oss://<bucket>/<object>` format. `<bucket>` specifies the name of the OSS bucket that is in the same region as the current project, and `<object>` specifies the path of the object with the extension included.
	//
	// example:
	//
	// oss://test-bucket/test-object2
	URI2 *string `json:"URI2,omitempty" xml:"URI2,omitempty"`
}

func (s CompareImageFacesRequestSource) String() string {
	return dara.Prettify(s)
}

func (s CompareImageFacesRequestSource) GoString() string {
	return s.String()
}

func (s *CompareImageFacesRequestSource) GetURI1() *string {
	return s.URI1
}

func (s *CompareImageFacesRequestSource) GetURI2() *string {
	return s.URI2
}

func (s *CompareImageFacesRequestSource) SetURI1(v string) *CompareImageFacesRequestSource {
	s.URI1 = &v
	return s
}

func (s *CompareImageFacesRequestSource) SetURI2(v string) *CompareImageFacesRequestSource {
	s.URI2 = &v
	return s
}

func (s *CompareImageFacesRequestSource) Validate() error {
	return dara.Validate(s)
}
