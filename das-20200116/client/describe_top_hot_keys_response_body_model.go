// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopHotKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTopHotKeysResponseBody
	GetCode() *string
	SetData(v *DescribeTopHotKeysResponseBodyData) *DescribeTopHotKeysResponseBody
	GetData() *DescribeTopHotKeysResponseBodyData
	SetMessage(v string) *DescribeTopHotKeysResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTopHotKeysResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeTopHotKeysResponseBody
	GetSuccess() *string
}

type DescribeTopHotKeysResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information about the hot keys.
	Data *DescribeTopHotKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTopHotKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopHotKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopHotKeysResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTopHotKeysResponseBody) GetData() *DescribeTopHotKeysResponseBodyData {
	return s.Data
}

func (s *DescribeTopHotKeysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTopHotKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTopHotKeysResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeTopHotKeysResponseBody) SetCode(v string) *DescribeTopHotKeysResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTopHotKeysResponseBody) SetData(v *DescribeTopHotKeysResponseBodyData) *DescribeTopHotKeysResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTopHotKeysResponseBody) SetMessage(v string) *DescribeTopHotKeysResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTopHotKeysResponseBody) SetRequestId(v string) *DescribeTopHotKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopHotKeysResponseBody) SetSuccess(v string) *DescribeTopHotKeysResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTopHotKeysResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTopHotKeysResponseBodyData struct {
	HotKey []*DescribeTopHotKeysResponseBodyDataHotKey `json:"HotKey,omitempty" xml:"HotKey,omitempty" type:"Repeated"`
}

func (s DescribeTopHotKeysResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopHotKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTopHotKeysResponseBodyData) GetHotKey() []*DescribeTopHotKeysResponseBodyDataHotKey {
	return s.HotKey
}

func (s *DescribeTopHotKeysResponseBodyData) SetHotKey(v []*DescribeTopHotKeysResponseBodyDataHotKey) *DescribeTopHotKeysResponseBodyData {
	s.HotKey = v
	return s
}

func (s *DescribeTopHotKeysResponseBodyData) Validate() error {
	if s.HotKey != nil {
		for _, item := range s.HotKey {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTopHotKeysResponseBodyDataHotKey struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The database in which the key is stored.
	//
	// example:
	//
	// 0
	Db *int32 `json:"Db,omitempty" xml:"Db,omitempty"`
	// The frequency at which the key is accessed, which indicates the QPS of the key.
	//
	// example:
	//
	// 5500~6000
	Hot     *string `json:"Hot,omitempty" xml:"Hot,omitempty"`
	InBytes *int64  `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The key.
	//
	// example:
	//
	// abc:def:eng
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The type of the key.
	//
	// example:
	//
	// zset
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// The statistical value that is calculated based on the least frequently used (LFU) caching algorithm.
	//
	// example:
	//
	// 253
	Lfu *int32 `json:"Lfu,omitempty" xml:"Lfu,omitempty"`
	// The ID of the data shard on the ApsaraDB for Redis instance.
	//
	// example:
	//
	// r-x****-db-0
	NodeId   *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OutBytes *int64  `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
}

func (s DescribeTopHotKeysResponseBodyDataHotKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopHotKeysResponseBodyDataHotKey) GoString() string {
	return s.String()
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) GetCategory() *string {
	return s.Category
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) GetDb() *int32 {
	return s.Db
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) GetHot() *string {
	return s.Hot
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) GetKey() *string {
	return s.Key
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) GetKeyType() *string {
	return s.KeyType
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) GetLfu() *int32 {
	return s.Lfu
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetCategory(v string) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.Category = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetDb(v int32) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.Db = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetHot(v string) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.Hot = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetInBytes(v int64) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.InBytes = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetKey(v string) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.Key = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetKeyType(v string) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.KeyType = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetLfu(v int32) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.Lfu = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetNodeId(v string) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.NodeId = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) SetOutBytes(v int64) *DescribeTopHotKeysResponseBodyDataHotKey {
	s.OutBytes = &v
	return s
}

func (s *DescribeTopHotKeysResponseBodyDataHotKey) Validate() error {
	return dara.Validate(s)
}
