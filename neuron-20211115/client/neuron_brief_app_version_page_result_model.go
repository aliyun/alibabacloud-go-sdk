// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronBriefAppVersionPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*NeuronBriefAppVersion) *NeuronBriefAppVersionPageResult
	GetData() []*NeuronBriefAppVersion
	SetRequestId(v string) *NeuronBriefAppVersionPageResult
	GetRequestId() *string
	SetTotal(v int32) *NeuronBriefAppVersionPageResult
	GetTotal() *int32
}

type NeuronBriefAppVersionPageResult struct {
	Data []*NeuronBriefAppVersion `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s NeuronBriefAppVersionPageResult) String() string {
	return dara.Prettify(s)
}

func (s NeuronBriefAppVersionPageResult) GoString() string {
	return s.String()
}

func (s *NeuronBriefAppVersionPageResult) GetData() []*NeuronBriefAppVersion {
	return s.Data
}

func (s *NeuronBriefAppVersionPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *NeuronBriefAppVersionPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *NeuronBriefAppVersionPageResult) SetData(v []*NeuronBriefAppVersion) *NeuronBriefAppVersionPageResult {
	s.Data = v
	return s
}

func (s *NeuronBriefAppVersionPageResult) SetRequestId(v string) *NeuronBriefAppVersionPageResult {
	s.RequestId = &v
	return s
}

func (s *NeuronBriefAppVersionPageResult) SetTotal(v int32) *NeuronBriefAppVersionPageResult {
	s.Total = &v
	return s
}

func (s *NeuronBriefAppVersionPageResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
