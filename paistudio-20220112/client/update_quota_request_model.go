// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateQuotaRequest
	GetDescription() *string
	SetLabels(v []*Label) *UpdateQuotaRequest
	GetLabels() []*Label
	SetQueueStrategy(v string) *UpdateQuotaRequest
	GetQueueStrategy() *string
	SetQuotaConfig(v *QuotaConfig) *UpdateQuotaRequest
	GetQuotaConfig() *QuotaConfig
	SetQuotaName(v string) *UpdateQuotaRequest
	GetQuotaName() *string
}

type UpdateQuotaRequest struct {
	// example:
	//
	// this is a test quota
	Description   *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels        []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	QueueStrategy *string  `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	// if can be null:
	// true
	QuotaConfig *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	QuotaName   *string      `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
}

func (s UpdateQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateQuotaRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *UpdateQuotaRequest) GetQueueStrategy() *string {
	return s.QueueStrategy
}

func (s *UpdateQuotaRequest) GetQuotaConfig() *QuotaConfig {
	return s.QuotaConfig
}

func (s *UpdateQuotaRequest) GetQuotaName() *string {
	return s.QuotaName
}

func (s *UpdateQuotaRequest) SetDescription(v string) *UpdateQuotaRequest {
	s.Description = &v
	return s
}

func (s *UpdateQuotaRequest) SetLabels(v []*Label) *UpdateQuotaRequest {
	s.Labels = v
	return s
}

func (s *UpdateQuotaRequest) SetQueueStrategy(v string) *UpdateQuotaRequest {
	s.QueueStrategy = &v
	return s
}

func (s *UpdateQuotaRequest) SetQuotaConfig(v *QuotaConfig) *UpdateQuotaRequest {
	s.QuotaConfig = v
	return s
}

func (s *UpdateQuotaRequest) SetQuotaName(v string) *UpdateQuotaRequest {
	s.QuotaName = &v
	return s
}

func (s *UpdateQuotaRequest) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QuotaConfig != nil {
		if err := s.QuotaConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
