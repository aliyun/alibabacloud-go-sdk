// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DescribeIndexResponseBody
	GetCollection() *string
	SetIndexDef(v string) *DescribeIndexResponseBody
	GetIndexDef() *string
	SetIndexName(v string) *DescribeIndexResponseBody
	GetIndexName() *string
	SetMessage(v string) *DescribeIndexResponseBody
	GetMessage() *string
	SetNamespace(v string) *DescribeIndexResponseBody
	GetNamespace() *string
	SetRequestId(v string) *DescribeIndexResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeIndexResponseBody
	GetStatus() *string
}

type DescribeIndexResponseBody struct {
	// The name of the collection.
	//
	// example:
	//
	// testcollection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The definition of the index.
	//
	// example:
	//
	// CREATE INDEX testindex ON mynamespace. testcollection
	IndexDef *string `json:"IndexDef,omitempty" xml:"IndexDef,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// testindex
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The namespace.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **false**: The operation fails.
	//
	// 	- **true**: The operation is successful.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIndexResponseBody) GetCollection() *string {
	return s.Collection
}

func (s *DescribeIndexResponseBody) GetIndexDef() *string {
	return s.IndexDef
}

func (s *DescribeIndexResponseBody) GetIndexName() *string {
	return s.IndexName
}

func (s *DescribeIndexResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeIndexResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIndexResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeIndexResponseBody) SetCollection(v string) *DescribeIndexResponseBody {
	s.Collection = &v
	return s
}

func (s *DescribeIndexResponseBody) SetIndexDef(v string) *DescribeIndexResponseBody {
	s.IndexDef = &v
	return s
}

func (s *DescribeIndexResponseBody) SetIndexName(v string) *DescribeIndexResponseBody {
	s.IndexName = &v
	return s
}

func (s *DescribeIndexResponseBody) SetMessage(v string) *DescribeIndexResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeIndexResponseBody) SetNamespace(v string) *DescribeIndexResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeIndexResponseBody) SetRequestId(v string) *DescribeIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIndexResponseBody) SetStatus(v string) *DescribeIndexResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
