// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSearchIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DropSearchIndexResponseBody
	GetCode() *string
	SetRequestId(v string) *DropSearchIndexResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DropSearchIndexResponseBody
	GetSuccess() *string
}

type DropSearchIndexResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
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

func (s DropSearchIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DropSearchIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DropSearchIndexResponseBody) GetCode() *string {
	return s.Code
}

func (s *DropSearchIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DropSearchIndexResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DropSearchIndexResponseBody) SetCode(v string) *DropSearchIndexResponseBody {
	s.Code = &v
	return s
}

func (s *DropSearchIndexResponseBody) SetRequestId(v string) *DropSearchIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DropSearchIndexResponseBody) SetSuccess(v string) *DropSearchIndexResponseBody {
	s.Success = &v
	return s
}

func (s *DropSearchIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
