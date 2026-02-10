// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveDomainModels(v *DescribeLiveDomainMappingResponseBodyLiveDomainModels) *DescribeLiveDomainMappingResponseBody
	GetLiveDomainModels() *DescribeLiveDomainMappingResponseBodyLiveDomainModels
	SetRequestId(v string) *DescribeLiveDomainMappingResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainMappingResponseBody struct {
	LiveDomainModels *DescribeLiveDomainMappingResponseBodyLiveDomainModels `json:"LiveDomainModels,omitempty" xml:"LiveDomainModels,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F6CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMappingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMappingResponseBody) GetLiveDomainModels() *DescribeLiveDomainMappingResponseBodyLiveDomainModels {
	return s.LiveDomainModels
}

func (s *DescribeLiveDomainMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainMappingResponseBody) SetLiveDomainModels(v *DescribeLiveDomainMappingResponseBodyLiveDomainModels) *DescribeLiveDomainMappingResponseBody {
	s.LiveDomainModels = v
	return s
}

func (s *DescribeLiveDomainMappingResponseBody) SetRequestId(v string) *DescribeLiveDomainMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainMappingResponseBody) Validate() error {
	if s.LiveDomainModels != nil {
		if err := s.LiveDomainModels.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainMappingResponseBodyLiveDomainModels struct {
	LiveDomainModel []*DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel `json:"LiveDomainModel,omitempty" xml:"LiveDomainModel,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainMappingResponseBodyLiveDomainModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMappingResponseBodyLiveDomainModels) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMappingResponseBodyLiveDomainModels) GetLiveDomainModel() []*DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel {
	return s.LiveDomainModel
}

func (s *DescribeLiveDomainMappingResponseBodyLiveDomainModels) SetLiveDomainModel(v []*DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel) *DescribeLiveDomainMappingResponseBodyLiveDomainModels {
	s.LiveDomainModel = v
	return s
}

func (s *DescribeLiveDomainMappingResponseBodyLiveDomainModels) Validate() error {
	if s.LiveDomainModel != nil {
		for _, item := range s.LiveDomainModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel) GetType() *string {
	return s.Type
}

func (s *DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel) SetDomainName(v string) *DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel) SetType(v string) *DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel {
	s.Type = &v
	return s
}

func (s *DescribeLiveDomainMappingResponseBodyLiveDomainModelsLiveDomainModel) Validate() error {
	return dara.Validate(s)
}
