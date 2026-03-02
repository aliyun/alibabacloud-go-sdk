// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppMobiPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*NeuronAppMobi) *NeuronAppMobiPageResult
	GetData() []*NeuronAppMobi
	SetRequestId(v string) *NeuronAppMobiPageResult
	GetRequestId() *string
	SetTotal(v int32) *NeuronAppMobiPageResult
	GetTotal() *int32
}

type NeuronAppMobiPageResult struct {
	Data []*NeuronAppMobi `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s NeuronAppMobiPageResult) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppMobiPageResult) GoString() string {
	return s.String()
}

func (s *NeuronAppMobiPageResult) GetData() []*NeuronAppMobi {
	return s.Data
}

func (s *NeuronAppMobiPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *NeuronAppMobiPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *NeuronAppMobiPageResult) SetData(v []*NeuronAppMobi) *NeuronAppMobiPageResult {
	s.Data = v
	return s
}

func (s *NeuronAppMobiPageResult) SetRequestId(v string) *NeuronAppMobiPageResult {
	s.RequestId = &v
	return s
}

func (s *NeuronAppMobiPageResult) SetTotal(v int32) *NeuronAppMobiPageResult {
	s.Total = &v
	return s
}

func (s *NeuronAppMobiPageResult) Validate() error {
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
