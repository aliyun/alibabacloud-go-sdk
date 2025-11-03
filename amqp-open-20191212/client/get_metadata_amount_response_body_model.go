// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetadataAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetadataAmountResponseBodyData) *GetMetadataAmountResponseBody
	GetData() *GetMetadataAmountResponseBodyData
	SetRequestId(v string) *GetMetadataAmountResponseBody
	GetRequestId() *string
}

type GetMetadataAmountResponseBody struct {
	// The returned data.
	Data *GetMetadataAmountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B75ACF23-2BEB-44AC-A0B6-AE14EDCA***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMetadataAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetadataAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponseBody) GetData() *GetMetadataAmountResponseBodyData {
	return s.Data
}

func (s *GetMetadataAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetadataAmountResponseBody) SetData(v *GetMetadataAmountResponseBodyData) *GetMetadataAmountResponseBody {
	s.Data = v
	return s
}

func (s *GetMetadataAmountResponseBody) SetRequestId(v string) *GetMetadataAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetadataAmountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMetadataAmountResponseBodyData struct {
	// The number of created exchanges on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 7
	CurrentExchanges *int32 `json:"CurrentExchanges,omitempty" xml:"CurrentExchanges,omitempty"`
	// The number of created queues on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 1
	CurrentQueues *int32 `json:"CurrentQueues,omitempty" xml:"CurrentQueues,omitempty"`
	// The number of created vhosts on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 1
	CurrentVirtualHosts *int32 `json:"CurrentVirtualHosts,omitempty" xml:"CurrentVirtualHosts,omitempty"`
	// The maximum number of exchanges that can be created on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 20
	MaxExchanges *int32 `json:"MaxExchanges,omitempty" xml:"MaxExchanges,omitempty"`
	// The maximum number of queues that can be created on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 20
	MaxQueues *int32 `json:"MaxQueues,omitempty" xml:"MaxQueues,omitempty"`
	// The maximum number of vhosts that can be created on the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// 10
	MaxVirtualHosts *int32 `json:"MaxVirtualHosts,omitempty" xml:"MaxVirtualHosts,omitempty"`
}

func (s GetMetadataAmountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetadataAmountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponseBodyData) GetCurrentExchanges() *int32 {
	return s.CurrentExchanges
}

func (s *GetMetadataAmountResponseBodyData) GetCurrentQueues() *int32 {
	return s.CurrentQueues
}

func (s *GetMetadataAmountResponseBodyData) GetCurrentVirtualHosts() *int32 {
	return s.CurrentVirtualHosts
}

func (s *GetMetadataAmountResponseBodyData) GetMaxExchanges() *int32 {
	return s.MaxExchanges
}

func (s *GetMetadataAmountResponseBodyData) GetMaxQueues() *int32 {
	return s.MaxQueues
}

func (s *GetMetadataAmountResponseBodyData) GetMaxVirtualHosts() *int32 {
	return s.MaxVirtualHosts
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentExchanges(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentExchanges = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentQueues(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentQueues = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetCurrentVirtualHosts(v int32) *GetMetadataAmountResponseBodyData {
	s.CurrentVirtualHosts = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxExchanges(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxExchanges = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxQueues(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxQueues = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) SetMaxVirtualHosts(v int32) *GetMetadataAmountResponseBodyData {
	s.MaxVirtualHosts = &v
	return s
}

func (s *GetMetadataAmountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
