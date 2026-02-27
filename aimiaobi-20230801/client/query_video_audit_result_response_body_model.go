// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoAuditResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryVideoAuditResultResponseBody
	GetCode() *string
	SetData(v *QueryVideoAuditResultResponseBodyData) *QueryVideoAuditResultResponseBody
	GetData() *QueryVideoAuditResultResponseBodyData
	SetHttpStatusCode(v int32) *QueryVideoAuditResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryVideoAuditResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryVideoAuditResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryVideoAuditResultResponseBody
	GetSuccess() *bool
}

type QueryVideoAuditResultResponseBody struct {
	// 业务处理结果状态码
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 视频审校的详细结果
	Data *QueryVideoAuditResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP响应状态码
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 业务处理结果描述信息
	//
	// example:
	//
	// 查询成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 本次API请求的唯一标识
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否处理成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryVideoAuditResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoAuditResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVideoAuditResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryVideoAuditResultResponseBody) GetData() *QueryVideoAuditResultResponseBodyData {
	return s.Data
}

func (s *QueryVideoAuditResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryVideoAuditResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryVideoAuditResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVideoAuditResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryVideoAuditResultResponseBody) SetCode(v string) *QueryVideoAuditResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVideoAuditResultResponseBody) SetData(v *QueryVideoAuditResultResponseBodyData) *QueryVideoAuditResultResponseBody {
	s.Data = v
	return s
}

func (s *QueryVideoAuditResultResponseBody) SetHttpStatusCode(v int32) *QueryVideoAuditResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryVideoAuditResultResponseBody) SetMessage(v string) *QueryVideoAuditResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryVideoAuditResultResponseBody) SetRequestId(v string) *QueryVideoAuditResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVideoAuditResultResponseBody) SetSuccess(v bool) *QueryVideoAuditResultResponseBody {
	s.Success = &v
	return s
}

func (s *QueryVideoAuditResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryVideoAuditResultResponseBodyData struct {
	// 视频总时长（秒）
	//
	// example:
	//
	// 120.5
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 任务执行失败时的错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 视频帧率（FPS）
	//
	// example:
	//
	// 30.0
	Fps *float64 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// 已经完成审核的帧数
	//
	// example:
	//
	// 120
	FrameAudited *int32 `json:"FrameAudited,omitempty" xml:"FrameAudited,omitempty"`
	// 视频高度（像素）
	//
	// example:
	//
	// 1080
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// 抽取的图片URL列表
	ImageUrls []*QueryVideoAuditResultResponseBodyDataImageUrls `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	// 图片审核结果详情
	Results []*QueryVideoAuditResultResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// 任务状态：PENDING(待执行)、RUNNING(执行中)、SUCCESSED(成功)、FAILED(失败)、CANCELED(取消)
	//
	// example:
	//
	// SUCCESSED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 视频审校的文本结果
	//
	// example:
	//
	// 视频审核完成
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 需要审核的视频帧总数
	//
	// example:
	//
	// 120
	TotalFrameAudit *int32 `json:"TotalFrameAudit,omitempty" xml:"TotalFrameAudit,omitempty"`
	// 视频总帧数
	//
	// example:
	//
	// 3615
	TotalFrames *int32 `json:"TotalFrames,omitempty" xml:"TotalFrames,omitempty"`
	// 检测到的视频分镜总数
	//
	// example:
	//
	// 15
	TotalShots *int32 `json:"TotalShots,omitempty" xml:"TotalShots,omitempty"`
	// 被审核的视频文件Key
	//
	// example:
	//
	// video/test.mp4
	VideoFileKey *string `json:"VideoFileKey,omitempty" xml:"VideoFileKey,omitempty"`
	// 被审核的视频URL地址
	//
	// example:
	//
	// https://example.com/video.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// 视频宽度（像素）
	//
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryVideoAuditResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoAuditResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryVideoAuditResultResponseBodyData) GetDuration() *float64 {
	return s.Duration
}

func (s *QueryVideoAuditResultResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryVideoAuditResultResponseBodyData) GetFps() *float64 {
	return s.Fps
}

func (s *QueryVideoAuditResultResponseBodyData) GetFrameAudited() *int32 {
	return s.FrameAudited
}

func (s *QueryVideoAuditResultResponseBodyData) GetHeight() *int32 {
	return s.Height
}

func (s *QueryVideoAuditResultResponseBodyData) GetImageUrls() []*QueryVideoAuditResultResponseBodyDataImageUrls {
	return s.ImageUrls
}

func (s *QueryVideoAuditResultResponseBodyData) GetResults() []*QueryVideoAuditResultResponseBodyDataResults {
	return s.Results
}

func (s *QueryVideoAuditResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryVideoAuditResultResponseBodyData) GetText() *string {
	return s.Text
}

func (s *QueryVideoAuditResultResponseBodyData) GetTotalFrameAudit() *int32 {
	return s.TotalFrameAudit
}

func (s *QueryVideoAuditResultResponseBodyData) GetTotalFrames() *int32 {
	return s.TotalFrames
}

func (s *QueryVideoAuditResultResponseBodyData) GetTotalShots() *int32 {
	return s.TotalShots
}

func (s *QueryVideoAuditResultResponseBodyData) GetVideoFileKey() *string {
	return s.VideoFileKey
}

func (s *QueryVideoAuditResultResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *QueryVideoAuditResultResponseBodyData) GetWidth() *int32 {
	return s.Width
}

func (s *QueryVideoAuditResultResponseBodyData) SetDuration(v float64) *QueryVideoAuditResultResponseBodyData {
	s.Duration = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetErrorMessage(v string) *QueryVideoAuditResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetFps(v float64) *QueryVideoAuditResultResponseBodyData {
	s.Fps = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetFrameAudited(v int32) *QueryVideoAuditResultResponseBodyData {
	s.FrameAudited = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetHeight(v int32) *QueryVideoAuditResultResponseBodyData {
	s.Height = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetImageUrls(v []*QueryVideoAuditResultResponseBodyDataImageUrls) *QueryVideoAuditResultResponseBodyData {
	s.ImageUrls = v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetResults(v []*QueryVideoAuditResultResponseBodyDataResults) *QueryVideoAuditResultResponseBodyData {
	s.Results = v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetStatus(v string) *QueryVideoAuditResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetText(v string) *QueryVideoAuditResultResponseBodyData {
	s.Text = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetTotalFrameAudit(v int32) *QueryVideoAuditResultResponseBodyData {
	s.TotalFrameAudit = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetTotalFrames(v int32) *QueryVideoAuditResultResponseBodyData {
	s.TotalFrames = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetTotalShots(v int32) *QueryVideoAuditResultResponseBodyData {
	s.TotalShots = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetVideoFileKey(v string) *QueryVideoAuditResultResponseBodyData {
	s.VideoFileKey = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetVideoUrl(v string) *QueryVideoAuditResultResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) SetWidth(v int32) *QueryVideoAuditResultResponseBodyData {
	s.Width = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyData) Validate() error {
	if s.ImageUrls != nil {
		for _, item := range s.ImageUrls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type QueryVideoAuditResultResponseBodyDataImageUrls struct {
	// 图片ID，与AliyunImageAuditResult中的dataId对应
	//
	// example:
	//
	// img001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1000
	Timestamp *float64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// https://example.com/image1.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryVideoAuditResultResponseBodyDataImageUrls) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoAuditResultResponseBodyDataImageUrls) GoString() string {
	return s.String()
}

func (s *QueryVideoAuditResultResponseBodyDataImageUrls) GetId() *string {
	return s.Id
}

func (s *QueryVideoAuditResultResponseBodyDataImageUrls) GetTimestamp() *float64 {
	return s.Timestamp
}

func (s *QueryVideoAuditResultResponseBodyDataImageUrls) GetUrl() *string {
	return s.Url
}

func (s *QueryVideoAuditResultResponseBodyDataImageUrls) SetId(v string) *QueryVideoAuditResultResponseBodyDataImageUrls {
	s.Id = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyDataImageUrls) SetTimestamp(v float64) *QueryVideoAuditResultResponseBodyDataImageUrls {
	s.Timestamp = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyDataImageUrls) SetUrl(v string) *QueryVideoAuditResultResponseBodyDataImageUrls {
	s.Url = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyDataImageUrls) Validate() error {
	return dara.Validate(s)
}

type QueryVideoAuditResultResponseBodyDataResults struct {
	// 对应图片的ID，与ImageUrl中的id字段对应
	//
	// example:
	//
	// d411ed15e8fc154fd0ef5addabfee04b
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// 审核请求ID
	//
	// example:
	//
	// B5D1CF9E-0404-51E3-A28E-A5C7D95B6C71
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	// 图片检测的风险标签、置信分等参数结果
	Result []*QueryVideoAuditResultResponseBodyDataResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// 风险等级：high(高风险)、medium(中风险)、low(低风险)、none(未检测到风险)
	//
	// example:
	//
	// none
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s QueryVideoAuditResultResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoAuditResultResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *QueryVideoAuditResultResponseBodyDataResults) GetDataId() *string {
	return s.DataId
}

func (s *QueryVideoAuditResultResponseBodyDataResults) GetReqId() *string {
	return s.ReqId
}

func (s *QueryVideoAuditResultResponseBodyDataResults) GetResult() []*QueryVideoAuditResultResponseBodyDataResultsResult {
	return s.Result
}

func (s *QueryVideoAuditResultResponseBodyDataResults) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *QueryVideoAuditResultResponseBodyDataResults) SetDataId(v string) *QueryVideoAuditResultResponseBodyDataResults {
	s.DataId = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyDataResults) SetReqId(v string) *QueryVideoAuditResultResponseBodyDataResults {
	s.ReqId = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyDataResults) SetResult(v []*QueryVideoAuditResultResponseBodyDataResultsResult) *QueryVideoAuditResultResponseBodyDataResults {
	s.Result = v
	return s
}

func (s *QueryVideoAuditResultResponseBodyDataResults) SetRiskLevel(v string) *QueryVideoAuditResultResponseBodyDataResults {
	s.RiskLevel = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyDataResults) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryVideoAuditResultResponseBodyDataResultsResult struct {
	// 0到100分，保留到小数点后2位，部分标签无置信分
	//
	// example:
	//
	// 99.5
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Label字段的解释说明
	//
	// example:
	//
	// 未检测出风险
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 图片内容检测运算后返回的标签，如：nonLabel（未检测出风险）
	//
	// example:
	//
	// nonLabel
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s QueryVideoAuditResultResponseBodyDataResultsResult) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoAuditResultResponseBodyDataResultsResult) GoString() string {
	return s.String()
}

func (s *QueryVideoAuditResultResponseBodyDataResultsResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *QueryVideoAuditResultResponseBodyDataResultsResult) GetDescription() *string {
	return s.Description
}

func (s *QueryVideoAuditResultResponseBodyDataResultsResult) GetLabel() *string {
	return s.Label
}

func (s *QueryVideoAuditResultResponseBodyDataResultsResult) SetConfidence(v float32) *QueryVideoAuditResultResponseBodyDataResultsResult {
	s.Confidence = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyDataResultsResult) SetDescription(v string) *QueryVideoAuditResultResponseBodyDataResultsResult {
	s.Description = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyDataResultsResult) SetLabel(v string) *QueryVideoAuditResultResponseBodyDataResultsResult {
	s.Label = &v
	return s
}

func (s *QueryVideoAuditResultResponseBodyDataResultsResult) Validate() error {
	return dara.Validate(s)
}
