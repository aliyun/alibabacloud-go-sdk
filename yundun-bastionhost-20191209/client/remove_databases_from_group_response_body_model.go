// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDatabasesFromGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveDatabasesFromGroupResponseBody
	GetRequestId() *string
	SetResults(v []*RemoveDatabasesFromGroupResponseBodyResults) *RemoveDatabasesFromGroupResponseBody
	GetResults() []*RemoveDatabasesFromGroupResponseBodyResults
}

type RemoveDatabasesFromGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*RemoveDatabasesFromGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RemoveDatabasesFromGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDatabasesFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDatabasesFromGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDatabasesFromGroupResponseBody) GetResults() []*RemoveDatabasesFromGroupResponseBodyResults {
	return s.Results
}

func (s *RemoveDatabasesFromGroupResponseBody) SetRequestId(v string) *RemoveDatabasesFromGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDatabasesFromGroupResponseBody) SetResults(v []*RemoveDatabasesFromGroupResponseBodyResults) *RemoveDatabasesFromGroupResponseBody {
	s.Results = v
	return s
}

func (s *RemoveDatabasesFromGroupResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RemoveDatabasesFromGroupResponseBodyResults struct {
	// The error code that is returned. If OK is returned, the operation was successful. If another error code is returned, the operation failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The database ID.
	//
	// example:
	//
	// 20
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The asset group ID.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The error message that is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveDatabasesFromGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s RemoveDatabasesFromGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RemoveDatabasesFromGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *RemoveDatabasesFromGroupResponseBodyResults) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *RemoveDatabasesFromGroupResponseBodyResults) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *RemoveDatabasesFromGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *RemoveDatabasesFromGroupResponseBodyResults) SetCode(v string) *RemoveDatabasesFromGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *RemoveDatabasesFromGroupResponseBodyResults) SetDatabaseId(v string) *RemoveDatabasesFromGroupResponseBodyResults {
	s.DatabaseId = &v
	return s
}

func (s *RemoveDatabasesFromGroupResponseBodyResults) SetHostGroupId(v string) *RemoveDatabasesFromGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *RemoveDatabasesFromGroupResponseBodyResults) SetMessage(v string) *RemoveDatabasesFromGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *RemoveDatabasesFromGroupResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
