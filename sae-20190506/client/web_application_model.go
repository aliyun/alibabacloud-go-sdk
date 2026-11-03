// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplication interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *WebApplication
	GetApplicationId() *string
	SetApplicationName(v string) *WebApplication
	GetApplicationName() *string
	SetCreatedTime(v string) *WebApplication
	GetCreatedTime() *string
	SetDescription(v string) *WebApplication
	GetDescription() *string
	SetInternetURL(v string) *WebApplication
	GetInternetURL() *string
	SetIntranetURL(v string) *WebApplication
	GetIntranetURL() *string
	SetLastModifiedTime(v string) *WebApplication
	GetLastModifiedTime() *string
	SetNamespaceId(v string) *WebApplication
	GetNamespaceId() *string
	SetRevisionConfig(v *RevisionConfig) *WebApplication
	GetRevisionConfig() *RevisionConfig
	SetVpcId(v string) *WebApplication
	GetVpcId() *string
	SetWebScalingConfig(v *WebScalingConfig) *WebApplication
	GetWebScalingConfig() *WebScalingConfig
	SetWebTrafficConfig(v *WebTrafficConfig) *WebApplication
	GetWebTrafficConfig() *WebTrafficConfig
}

type WebApplication struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// sae-app
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2019-07-11 15:54:49
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The application description.
	//
	// example:
	//
	// my sae app
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The public URL for accessing the application.
	//
	// example:
	//
	// https://****.cn-hangzhou.sae.run
	InternetURL *string `json:"InternetURL,omitempty" xml:"InternetURL,omitempty"`
	// The internal URL for accessing the application.
	//
	// example:
	//
	// https://****.cn-hangzhou-vpc.sae.run
	IntranetURL *string `json:"IntranetURL,omitempty" xml:"IntranetURL,omitempty"`
	// The time when the application was last modified.
	//
	// example:
	//
	// 2019-07-11 20:12:58
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The ID of the namespace to which the application belongs.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The configurations of the revision version.
	RevisionConfig *RevisionConfig `json:"RevisionConfig,omitempty" xml:"RevisionConfig,omitempty"`
	// The Virtual Private Cloud (VPC) ID of the application namespace.
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The scaling configurations of the application.
	WebScalingConfig *WebScalingConfig `json:"WebScalingConfig,omitempty" xml:"WebScalingConfig,omitempty"`
	// The traffic configurations of the application.
	WebTrafficConfig *WebTrafficConfig `json:"WebTrafficConfig,omitempty" xml:"WebTrafficConfig,omitempty"`
}

func (s WebApplication) String() string {
	return dara.Prettify(s)
}

func (s WebApplication) GoString() string {
	return s.String()
}

func (s *WebApplication) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *WebApplication) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *WebApplication) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *WebApplication) GetDescription() *string {
	return s.Description
}

func (s *WebApplication) GetInternetURL() *string {
	return s.InternetURL
}

func (s *WebApplication) GetIntranetURL() *string {
	return s.IntranetURL
}

func (s *WebApplication) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *WebApplication) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *WebApplication) GetRevisionConfig() *RevisionConfig {
	return s.RevisionConfig
}

func (s *WebApplication) GetVpcId() *string {
	return s.VpcId
}

func (s *WebApplication) GetWebScalingConfig() *WebScalingConfig {
	return s.WebScalingConfig
}

func (s *WebApplication) GetWebTrafficConfig() *WebTrafficConfig {
	return s.WebTrafficConfig
}

func (s *WebApplication) SetApplicationId(v string) *WebApplication {
	s.ApplicationId = &v
	return s
}

func (s *WebApplication) SetApplicationName(v string) *WebApplication {
	s.ApplicationName = &v
	return s
}

func (s *WebApplication) SetCreatedTime(v string) *WebApplication {
	s.CreatedTime = &v
	return s
}

func (s *WebApplication) SetDescription(v string) *WebApplication {
	s.Description = &v
	return s
}

func (s *WebApplication) SetInternetURL(v string) *WebApplication {
	s.InternetURL = &v
	return s
}

func (s *WebApplication) SetIntranetURL(v string) *WebApplication {
	s.IntranetURL = &v
	return s
}

func (s *WebApplication) SetLastModifiedTime(v string) *WebApplication {
	s.LastModifiedTime = &v
	return s
}

func (s *WebApplication) SetNamespaceId(v string) *WebApplication {
	s.NamespaceId = &v
	return s
}

func (s *WebApplication) SetRevisionConfig(v *RevisionConfig) *WebApplication {
	s.RevisionConfig = v
	return s
}

func (s *WebApplication) SetVpcId(v string) *WebApplication {
	s.VpcId = &v
	return s
}

func (s *WebApplication) SetWebScalingConfig(v *WebScalingConfig) *WebApplication {
	s.WebScalingConfig = v
	return s
}

func (s *WebApplication) SetWebTrafficConfig(v *WebTrafficConfig) *WebApplication {
	s.WebTrafficConfig = v
	return s
}

func (s *WebApplication) Validate() error {
	if s.RevisionConfig != nil {
		if err := s.RevisionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WebScalingConfig != nil {
		if err := s.WebScalingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WebTrafficConfig != nil {
		if err := s.WebTrafficConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
