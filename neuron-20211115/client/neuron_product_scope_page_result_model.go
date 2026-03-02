// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronProductScopePageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*NeuronProductScope) *NeuronProductScopePageResult
	GetData() []*NeuronProductScope
	SetRequestId(v string) *NeuronProductScopePageResult
	GetRequestId() *string
	SetTotal(v int32) *NeuronProductScopePageResult
	GetTotal() *int32
}

type NeuronProductScopePageResult struct {
	Data []*NeuronProductScope `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s NeuronProductScopePageResult) String() string {
	return dara.Prettify(s)
}

func (s NeuronProductScopePageResult) GoString() string {
	return s.String()
}

func (s *NeuronProductScopePageResult) GetData() []*NeuronProductScope {
	return s.Data
}

func (s *NeuronProductScopePageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *NeuronProductScopePageResult) GetTotal() *int32 {
	return s.Total
}

func (s *NeuronProductScopePageResult) SetData(v []*NeuronProductScope) *NeuronProductScopePageResult {
	s.Data = v
	return s
}

func (s *NeuronProductScopePageResult) SetRequestId(v string) *NeuronProductScopePageResult {
	s.RequestId = &v
	return s
}

func (s *NeuronProductScopePageResult) SetTotal(v int32) *NeuronProductScopePageResult {
	s.Total = &v
	return s
}

func (s *NeuronProductScopePageResult) Validate() error {
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
