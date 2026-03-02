// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronClientInfoPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*NeuronAppClientInfo) *NeuronClientInfoPageResult
	GetData() []*NeuronAppClientInfo
	SetRequestId(v string) *NeuronClientInfoPageResult
	GetRequestId() *string
	SetShowApp(v bool) *NeuronClientInfoPageResult
	GetShowApp() *bool
	SetTotal(v int32) *NeuronClientInfoPageResult
	GetTotal() *int32
}

type NeuronClientInfoPageResult struct {
	Data []*NeuronAppClientInfo `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	ShowApp *bool `json:"showApp,omitempty" xml:"showApp,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s NeuronClientInfoPageResult) String() string {
	return dara.Prettify(s)
}

func (s NeuronClientInfoPageResult) GoString() string {
	return s.String()
}

func (s *NeuronClientInfoPageResult) GetData() []*NeuronAppClientInfo {
	return s.Data
}

func (s *NeuronClientInfoPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *NeuronClientInfoPageResult) GetShowApp() *bool {
	return s.ShowApp
}

func (s *NeuronClientInfoPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *NeuronClientInfoPageResult) SetData(v []*NeuronAppClientInfo) *NeuronClientInfoPageResult {
	s.Data = v
	return s
}

func (s *NeuronClientInfoPageResult) SetRequestId(v string) *NeuronClientInfoPageResult {
	s.RequestId = &v
	return s
}

func (s *NeuronClientInfoPageResult) SetShowApp(v bool) *NeuronClientInfoPageResult {
	s.ShowApp = &v
	return s
}

func (s *NeuronClientInfoPageResult) SetTotal(v int32) *NeuronClientInfoPageResult {
	s.Total = &v
	return s
}

func (s *NeuronClientInfoPageResult) Validate() error {
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
