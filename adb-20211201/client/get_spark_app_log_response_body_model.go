// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSparkAppLogResponseBodyData) *GetSparkAppLogResponseBody
	GetData() *GetSparkAppLogResponseBodyData
	SetRequestId(v string) *GetSparkAppLogResponseBody
	GetRequestId() *string
}

type GetSparkAppLogResponseBody struct {
	// The queried log.
	Data *GetSparkAppLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C3A9594F-1D40-4472-A96C-8FB8AA20D38C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppLogResponseBody) GetData() *GetSparkAppLogResponseBodyData {
	return s.Data
}

func (s *GetSparkAppLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkAppLogResponseBody) SetData(v *GetSparkAppLogResponseBodyData) *GetSparkAppLogResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppLogResponseBody) SetRequestId(v string) *GetSparkAppLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkAppLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSparkAppLogResponseBodyData struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-clusterxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The content of the log.
	//
	// example:
	//
	// 22/04/22 15:30:49 INFO Utils: Start the dump task because s202206061441hz22a35ab000****-0001 app end, the interval is 238141ms;22/04/22 15:30:49 INFO AbstractConnector: Stopped Spark@5e774d9d{HTTP/1.1, (http/1.1)}{0.0.0.0:4040}
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	// The number of log entries. A value of 0 indicates that no valid logs are returned.
	//
	// example:
	//
	// 3517972480
	LogSize *int32 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// The alert message returned for the request, such as task execution failure or insufficient resources. If no alert occurs, null is returned.
	//
	// example:
	//
	// WARNING:  log file maybe deleted, please check oss path: oss://TestBucketName/applog/
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetSparkAppLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkAppLogResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkAppLogResponseBodyData) GetLogContent() *string {
	return s.LogContent
}

func (s *GetSparkAppLogResponseBodyData) GetLogSize() *int32 {
	return s.LogSize
}

func (s *GetSparkAppLogResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetSparkAppLogResponseBodyData) SetDBClusterId(v string) *GetSparkAppLogResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppLogResponseBodyData) SetLogContent(v string) *GetSparkAppLogResponseBodyData {
	s.LogContent = &v
	return s
}

func (s *GetSparkAppLogResponseBodyData) SetLogSize(v int32) *GetSparkAppLogResponseBodyData {
	s.LogSize = &v
	return s
}

func (s *GetSparkAppLogResponseBodyData) SetMessage(v string) *GetSparkAppLogResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetSparkAppLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
