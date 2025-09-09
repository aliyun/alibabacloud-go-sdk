// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDatabasesToGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDatabasesToGroupResponseBody
	GetRequestId() *string
	SetResults(v []*AddDatabasesToGroupResponseBodyResults) *AddDatabasesToGroupResponseBody
	GetResults() []*AddDatabasesToGroupResponseBodyResults
}

type AddDatabasesToGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned.
	Results []*AddDatabasesToGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AddDatabasesToGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDatabasesToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddDatabasesToGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDatabasesToGroupResponseBody) GetResults() []*AddDatabasesToGroupResponseBodyResults {
	return s.Results
}

func (s *AddDatabasesToGroupResponseBody) SetRequestId(v string) *AddDatabasesToGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDatabasesToGroupResponseBody) SetResults(v []*AddDatabasesToGroupResponseBodyResults) *AddDatabasesToGroupResponseBody {
	s.Results = v
	return s
}

func (s *AddDatabasesToGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddDatabasesToGroupResponseBodyResults struct {
	// The error code returned. If **OK*	- is returned, the operation was successful. If another error code is returned, the operation failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The database ID.
	//
	// example:
	//
	// 9
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The asset group ID.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AddDatabasesToGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AddDatabasesToGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AddDatabasesToGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AddDatabasesToGroupResponseBodyResults) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *AddDatabasesToGroupResponseBodyResults) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *AddDatabasesToGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AddDatabasesToGroupResponseBodyResults) SetCode(v string) *AddDatabasesToGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AddDatabasesToGroupResponseBodyResults) SetDatabaseId(v string) *AddDatabasesToGroupResponseBodyResults {
	s.DatabaseId = &v
	return s
}

func (s *AddDatabasesToGroupResponseBodyResults) SetHostGroupId(v string) *AddDatabasesToGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *AddDatabasesToGroupResponseBodyResults) SetMessage(v string) *AddDatabasesToGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AddDatabasesToGroupResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
