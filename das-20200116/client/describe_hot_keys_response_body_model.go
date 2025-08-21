// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHotKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHotKeysResponseBody
	GetCode() *string
	SetData(v *DescribeHotKeysResponseBodyData) *DescribeHotKeysResponseBody
	GetData() *DescribeHotKeysResponseBodyData
	SetMessage(v string) *DescribeHotKeysResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHotKeysResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeHotKeysResponseBody
	GetSuccess() *string
}

type DescribeHotKeysResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the hot keys.
	Data *DescribeHotKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
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
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHotKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHotKeysResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHotKeysResponseBody) GetData() *DescribeHotKeysResponseBodyData {
	return s.Data
}

func (s *DescribeHotKeysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHotKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHotKeysResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeHotKeysResponseBody) SetCode(v string) *DescribeHotKeysResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHotKeysResponseBody) SetData(v *DescribeHotKeysResponseBodyData) *DescribeHotKeysResponseBody {
	s.Data = v
	return s
}

func (s *DescribeHotKeysResponseBody) SetMessage(v string) *DescribeHotKeysResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHotKeysResponseBody) SetRequestId(v string) *DescribeHotKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHotKeysResponseBody) SetSuccess(v string) *DescribeHotKeysResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHotKeysResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHotKeysResponseBodyData struct {
	HotKey []*DescribeHotKeysResponseBodyDataHotKey `json:"HotKey,omitempty" xml:"HotKey,omitempty" type:"Repeated"`
}

func (s DescribeHotKeysResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeHotKeysResponseBodyData) GetHotKey() []*DescribeHotKeysResponseBodyDataHotKey {
	return s.HotKey
}

func (s *DescribeHotKeysResponseBodyData) SetHotKey(v []*DescribeHotKeysResponseBodyDataHotKey) *DescribeHotKeysResponseBodyData {
	s.HotKey = v
	return s
}

func (s *DescribeHotKeysResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeHotKeysResponseBodyDataHotKey struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The database in which the key is stored.
	//
	// example:
	//
	// 0
	Db *int32 `json:"Db,omitempty" xml:"Db,omitempty"`
	// The frequency at which the key is accessed, which indicates the queries per second (QPS) of the key.
	//
	// example:
	//
	// 5500~6000
	Hot     *string `json:"Hot,omitempty" xml:"Hot,omitempty"`
	InBytes *int64  `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The name of the key.
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
	KeyType  *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	NodeId   *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OutBytes *int64  `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The number of elements in the key.
	//
	// example:
	//
	// 2
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeHotKeysResponseBodyDataHotKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotKeysResponseBodyDataHotKey) GoString() string {
	return s.String()
}

func (s *DescribeHotKeysResponseBodyDataHotKey) GetCategory() *string {
	return s.Category
}

func (s *DescribeHotKeysResponseBodyDataHotKey) GetDb() *int32 {
	return s.Db
}

func (s *DescribeHotKeysResponseBodyDataHotKey) GetHot() *string {
	return s.Hot
}

func (s *DescribeHotKeysResponseBodyDataHotKey) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeHotKeysResponseBodyDataHotKey) GetKey() *string {
	return s.Key
}

func (s *DescribeHotKeysResponseBodyDataHotKey) GetKeyType() *string {
	return s.KeyType
}

func (s *DescribeHotKeysResponseBodyDataHotKey) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeHotKeysResponseBodyDataHotKey) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeHotKeysResponseBodyDataHotKey) GetSize() *int64 {
	return s.Size
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetCategory(v string) *DescribeHotKeysResponseBodyDataHotKey {
	s.Category = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetDb(v int32) *DescribeHotKeysResponseBodyDataHotKey {
	s.Db = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetHot(v string) *DescribeHotKeysResponseBodyDataHotKey {
	s.Hot = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetInBytes(v int64) *DescribeHotKeysResponseBodyDataHotKey {
	s.InBytes = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetKey(v string) *DescribeHotKeysResponseBodyDataHotKey {
	s.Key = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetKeyType(v string) *DescribeHotKeysResponseBodyDataHotKey {
	s.KeyType = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetNodeId(v string) *DescribeHotKeysResponseBodyDataHotKey {
	s.NodeId = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetOutBytes(v int64) *DescribeHotKeysResponseBodyDataHotKey {
	s.OutBytes = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) SetSize(v int64) *DescribeHotKeysResponseBodyDataHotKey {
	s.Size = &v
	return s
}

func (s *DescribeHotKeysResponseBodyDataHotKey) Validate() error {
	return dara.Validate(s)
}
