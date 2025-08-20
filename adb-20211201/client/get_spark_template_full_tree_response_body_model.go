// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkTemplateFullTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetSparkTemplateFullTreeResponseBody
	GetData() *string
	SetRequestId(v string) *GetSparkTemplateFullTreeResponseBody
	GetRequestId() *string
}

type GetSparkTemplateFullTreeResponseBody struct {
	// The directory structure of Spark applications. Fields in the response parameter:
	//
	// 	- **Uid**: the UID of the Alibaba Cloud account.
	//
	// 	- **Type**: the application template type. Valid values:
	//
	//     	- **FOLDER**
	//
	//     	- **FILE**
	//
	// 	- **Parent**: indicates whether a child directory exists. Valid values:
	//
	//     	- **0**: no.
	//
	//     	- **-1**: yes.
	//
	// 	- **Children**: the child directory.
	//
	// 	- **LastModified**: the time when applications are last modified. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// 	- **AppType**: the application type. Valid values:
	//
	//     	- **SQL**
	//
	//     	- **STREAMING**
	//
	//     	- **BATCH**
	//
	// 	- **Name**: the name of the directory or application.
	//
	// 	- **Id**: the directory ID or application ID.
	//
	// example:
	//
	// {     "Uid": 10415777****,     "Type": "FOLDER",     "Parent": -1,     "Children": [       {         "LastModified": 1648544748,         "Uid": 104157779****,         "Type": "FILE",         "Parent": 0,         "Id": s202204132****,         "AppType": "SQL",         "Name": "f"       },       {         "LastModified": 1648544956,         "Uid": 1041577795****,         "Type": "FOLDER",         "Parent": 0,         "Id": 157,         "Name": "f3333"       }     ],     "Id": 725204,     "Name": "root"   }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkTemplateFullTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkTemplateFullTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFullTreeResponseBody) GetData() *string {
	return s.Data
}

func (s *GetSparkTemplateFullTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkTemplateFullTreeResponseBody) SetData(v string) *GetSparkTemplateFullTreeResponseBody {
	s.Data = &v
	return s
}

func (s *GetSparkTemplateFullTreeResponseBody) SetRequestId(v string) *GetSparkTemplateFullTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkTemplateFullTreeResponseBody) Validate() error {
	return dara.Validate(s)
}
