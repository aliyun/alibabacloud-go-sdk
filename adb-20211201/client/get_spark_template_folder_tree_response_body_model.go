// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkTemplateFolderTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetSparkTemplateFolderTreeResponseBody
	GetData() *string
	SetRequestId(v string) *GetSparkTemplateFolderTreeResponseBody
	GetRequestId() *string
}

type GetSparkTemplateFolderTreeResponseBody struct {
	// The directory structure of Spark applications, which is in the tree format. Fields in the response parameter:
	//
	// 	- **Uid**: the UID of the Alibaba Cloud account.
	//
	// 	- **Type**: the application template type. Valid values: **FOLDER**
	//
	// 	- **Parent**: indicates whether a child directory exists. Valid values:
	//
	//     	- **0**: no.
	//
	//     	- **-1**: yes.
	//
	// 	- **Children**: the child directory.
	//
	// 	- **LastModified**: the time when applications in the directory are last modified. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// 	- **Name**: the name of the directory.
	//
	// 	- **Id**: the directory ID.
	//
	// example:
	//
	// {           "Uid":195813423****,           "Type":"FOLDER",          "Parent":-1,           "Children":[              {                     "LastModified":1647853173,               "Uid":195813423****,                     "Type":"FOLDER",                     "Parent":0,                     "Id":157,                     "Name":"t"         }       ],            "Id":725204,            "Name":"root"      }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkTemplateFolderTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkTemplateFolderTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFolderTreeResponseBody) GetData() *string {
	return s.Data
}

func (s *GetSparkTemplateFolderTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkTemplateFolderTreeResponseBody) SetData(v string) *GetSparkTemplateFolderTreeResponseBody {
	s.Data = &v
	return s
}

func (s *GetSparkTemplateFolderTreeResponseBody) SetRequestId(v string) *GetSparkTemplateFolderTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkTemplateFolderTreeResponseBody) Validate() error {
	return dara.Validate(s)
}
