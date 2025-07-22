// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopBigKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTopBigKeysResponseBody
	GetCode() *string
	SetData(v *DescribeTopBigKeysResponseBodyData) *DescribeTopBigKeysResponseBody
	GetData() *DescribeTopBigKeysResponseBodyData
	SetMessage(v string) *DescribeTopBigKeysResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTopBigKeysResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeTopBigKeysResponseBody
	GetSuccess() *string
}

type DescribeTopBigKeysResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information about the large keys.
	//
	// > This parameter is left empty If no large keys exist within the specified time range.
	Data *DescribeTopBigKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeTopBigKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopBigKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopBigKeysResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTopBigKeysResponseBody) GetData() *DescribeTopBigKeysResponseBodyData {
	return s.Data
}

func (s *DescribeTopBigKeysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTopBigKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTopBigKeysResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeTopBigKeysResponseBody) SetCode(v string) *DescribeTopBigKeysResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTopBigKeysResponseBody) SetData(v *DescribeTopBigKeysResponseBodyData) *DescribeTopBigKeysResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTopBigKeysResponseBody) SetMessage(v string) *DescribeTopBigKeysResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTopBigKeysResponseBody) SetRequestId(v string) *DescribeTopBigKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopBigKeysResponseBody) SetSuccess(v string) *DescribeTopBigKeysResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTopBigKeysResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTopBigKeysResponseBodyData struct {
	BigKey []*DescribeTopBigKeysResponseBodyDataBigKey `json:"BigKey,omitempty" xml:"BigKey,omitempty" type:"Repeated"`
}

func (s DescribeTopBigKeysResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopBigKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTopBigKeysResponseBodyData) GetBigKey() []*DescribeTopBigKeysResponseBodyDataBigKey {
	return s.BigKey
}

func (s *DescribeTopBigKeysResponseBodyData) SetBigKey(v []*DescribeTopBigKeysResponseBodyDataBigKey) *DescribeTopBigKeysResponseBodyData {
	s.BigKey = v
	return s
}

func (s *DescribeTopBigKeysResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeTopBigKeysResponseBodyDataBigKey struct {
	// The database in which the key is stored.
	//
	// example:
	//
	// 0
	Db *int32 `json:"Db,omitempty" xml:"Db,omitempty"`
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
	// The ID of the data shard on the ApsaraDB for Redis instance.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The number of elements in the key.
	//
	// example:
	//
	// 2
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeTopBigKeysResponseBodyDataBigKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopBigKeysResponseBodyDataBigKey) GoString() string {
	return s.String()
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) GetDb() *int32 {
	return s.Db
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) GetKey() *string {
	return s.Key
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) GetKeyType() *string {
	return s.KeyType
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) GetSize() *int64 {
	return s.Size
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) SetDb(v int32) *DescribeTopBigKeysResponseBodyDataBigKey {
	s.Db = &v
	return s
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) SetKey(v string) *DescribeTopBigKeysResponseBodyDataBigKey {
	s.Key = &v
	return s
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) SetKeyType(v string) *DescribeTopBigKeysResponseBodyDataBigKey {
	s.KeyType = &v
	return s
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) SetNodeId(v string) *DescribeTopBigKeysResponseBodyDataBigKey {
	s.NodeId = &v
	return s
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) SetSize(v int64) *DescribeTopBigKeysResponseBodyDataBigKey {
	s.Size = &v
	return s
}

func (s *DescribeTopBigKeysResponseBodyDataBigKey) Validate() error {
	return dara.Validate(s)
}
