// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*NeuronApp) *NeuronAppPageResult
	GetData() []*NeuronApp
	SetRequestId(v string) *NeuronAppPageResult
	GetRequestId() *string
	SetTotal(v int32) *NeuronAppPageResult
	GetTotal() *int32
}

type NeuronAppPageResult struct {
	Data []*NeuronApp `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s NeuronAppPageResult) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppPageResult) GoString() string {
	return s.String()
}

func (s *NeuronAppPageResult) GetData() []*NeuronApp {
	return s.Data
}

func (s *NeuronAppPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *NeuronAppPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *NeuronAppPageResult) SetData(v []*NeuronApp) *NeuronAppPageResult {
	s.Data = v
	return s
}

func (s *NeuronAppPageResult) SetRequestId(v string) *NeuronAppPageResult {
	s.RequestId = &v
	return s
}

func (s *NeuronAppPageResult) SetTotal(v int32) *NeuronAppPageResult {
	s.Total = &v
	return s
}

func (s *NeuronAppPageResult) Validate() error {
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
