// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogStore(v string) *GetLogMetaRequest
	GetLogStore() *string
	SetResourceDirectoryAccountId(v int64) *GetLogMetaRequest
	GetResourceDirectoryAccountId() *int64
}

type GetLogMetaRequest struct {
	// The name of the dedicated Logstore in which logs are stored.
	//
	// >  You can call the [DescribeLogMeta](~~DescribeLogMeta~~) operation to query the name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// aegis-log-login
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s GetLogMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogMetaRequest) GoString() string {
	return s.String()
}

func (s *GetLogMetaRequest) GetLogStore() *string {
	return s.LogStore
}

func (s *GetLogMetaRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *GetLogMetaRequest) SetLogStore(v string) *GetLogMetaRequest {
	s.LogStore = &v
	return s
}

func (s *GetLogMetaRequest) SetResourceDirectoryAccountId(v int64) *GetLogMetaRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetLogMetaRequest) Validate() error {
	return dara.Validate(s)
}
