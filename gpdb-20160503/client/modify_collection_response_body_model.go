// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifyCollectionResponseBody
	GetMessage() *string
	SetMetadata(v string) *ModifyCollectionResponseBody
	GetMetadata() *string
	SetRequestId(v string) *ModifyCollectionResponseBody
	GetRequestId() *string
	SetStatus(v string) *ModifyCollectionResponseBody
	GetStatus() *string
}

type ModifyCollectionResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The metadata of the vector data, which is a JSON string in the MAP format. The key specifies the field name, and the value specifies the data type.
	//
	// >
	//
	// 	- For information about the supported data types, see [Data types](https://help.aliyun.com/document_detail/424383.html).
	//
	// 	- The money data type is not supported.
	//
	// example:
	//
	// {"title":"text","content":"text","response":"int"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCollectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyCollectionResponseBody) GetMetadata() *string {
	return s.Metadata
}

func (s *ModifyCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCollectionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ModifyCollectionResponseBody) SetMessage(v string) *ModifyCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCollectionResponseBody) SetMetadata(v string) *ModifyCollectionResponseBody {
	s.Metadata = &v
	return s
}

func (s *ModifyCollectionResponseBody) SetRequestId(v string) *ModifyCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCollectionResponseBody) SetStatus(v string) *ModifyCollectionResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
